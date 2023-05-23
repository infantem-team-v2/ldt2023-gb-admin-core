package usecase

import (
	"context"
	"fmt"
	"gb-auth-gate/config"
	authInterface "gb-auth-gate/internal/auth/interface"
	"gb-auth-gate/internal/auth/model"
	"gb-auth-gate/internal/pkg/common"
	"gb-auth-gate/pkg/terrors"
	"gb-auth-gate/pkg/thttp/server"
	tlogger "gb-auth-gate/pkg/tlogger"
	"gb-auth-gate/pkg/tsecure"
	"gb-auth-gate/pkg/tutils/ptr"
	"github.com/fiorix/go-redis/redis"
	"github.com/sarulabs/di"
	"math/rand"
	"time"
)

type AuthUC struct {
	config   *config.Config
	logger   tlogger.ILogger
	authRepo authInterface.RelationalRepository
	fernet   *tsecure.FernetCrypto
	redis    *redis.Client
}

func BuildAuthUsecase(ctn di.Container) (interface{}, error) {
	return &AuthUC{
		config:   ctn.Get("config").(*config.Config),
		logger:   ctn.Get("logger").(tlogger.ILogger),
		authRepo: ctn.Get("authRepo").(authInterface.RelationalRepository),
		fernet:   ctn.Get("fernet").(*tsecure.FernetCrypto),
		redis:    ctn.Get("redis").(*redis.Client),
	}, nil
}

func (as *AuthUC) SignUp(params *model.SignUpRequest) (*model.SignUpResponse, error) {
	if params.BusinessData.INN == nil {
		return nil, terrors.Raise(nil, 100013)
	}
	if params.AuthData.Password != params.AuthData.RepeatedPassword {
		return nil, terrors.Raise(nil, 100014)
	}
	// TODO: add validation for inn by FNS api and getting company's name by same way
	if params.BusinessData.Name == nil {
		if len(*params.BusinessData.INN) == 12 {
			params.BusinessData.Name = ptr.String(fmt.Sprintf("ИП %s", *params.PersonalData.FullName))
		}

	}

	params.AuthData.Password = tsecure.CalcSignature(as.config.SecureConfig.Fernet.Key, params.AuthData.Password, tsecure.SHA512)

	userId, err := as.authRepo.CreateUser(context.Background(), &model.CreateUserDAO{
		FullName:         params.PersonalData.FullName,
		Email:            &params.AuthData.Email,
		Password:         &params.AuthData.Password,
		Inn:              params.BusinessData.INN,
		JobPosition:      params.PersonalData.Position,
		City:             params.PersonalData.Geographic.City,
		Country:          params.PersonalData.Geographic.Country,
		BusinessName:     params.BusinessData.Name,
		EconomicActivity: params.BusinessData.EconomicActivity,
		Website:          params.BusinessData.Website,
	})
	if err != nil {
		return nil, err
	}

	accessToken, refreshToken, err := as.GenerateTokensPair(&model.CreateAuthTokensLogic{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}

	return &model.SignUpResponse{
		AuthTokensLogic: model.AuthTokensLogic{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
		Response: common.Response{
			Message:      "Successfully created user",
			InternalCode: 201,
		},
	}, nil
}

func (as *AuthUC) ValidateEmail(params *model.EmailValidateRequest) (*model.EmailValidateResponse, error) {
	rand.Seed(time.Now().UnixNano())
	valid := rand.Intn(2) == 1
	return &model.EmailValidateResponse{
		Response: common.Response{
			Message:      "SUCCESS",
			InternalCode: 200,
		},
		Valid: valid,
	}, nil
}

func (as *AuthUC) SignIn(params *model.SignInRequest) (*model.SignInResponse, error) {
	user, err := as.authRepo.FindUserByEmail(params.Email)
	if err != nil {
		return nil, err
	}

	inputPassword := tsecure.CalcSignature(as.config.SecureConfig.Fernet.Key, params.Password, tsecure.SHA512)

	if user.Password != inputPassword {
		return nil, terrors.Raise(nil, 100016)
	}

	accessToken, refreshToken, err := as.GenerateTokensPair(&model.CreateAuthTokensLogic{
		UserId: user.UserId,
	})

	return &model.SignInResponse{
		AuthTokensLogic: model.AuthTokensLogic{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
		Response: common.Response{
			Message:      "Successful login",
			InternalCode: 200,
		},
	}, nil
}

func (as *AuthUC) ValidateService(params *model.AuthHeadersLogic) (bool, error) {
	service, err := as.authRepo.FindServiceByName(params.PublicKey)
	if err != nil {
		return false, err
	}
	decryptedPrivateKey, err := as.fernet.Decrypt(service.PrivateKey)
	if err != nil {
		return false, terrors.Raise(err, 200003)
	}
	signature := tsecure.CalcSignature(
		decryptedPrivateKey,
		string(params.Body),
		tsecure.SHA512,
	)
	if params.Signature == signature {
		return true, nil
	}

	return false, terrors.Raise(nil, 100003)
}

func (as *AuthUC) GenerateAccessToken(refreshToken string, params *model.CreateAuthTokensLogic) (accessToken string, err error) {
	duration := time.Now().Add(time.Second * time.Duration(as.config.HttpConfig.AccessExpireTime))
	accessToken, err = server.CreateJwtToken(&server.JwtParams{
		Salt:     &as.config.HttpConfig.JWTSalt,
		UserId:   &params.UserId,
		Duration: &duration,
		Type:     ptr.String("access"),
	})

	if err != nil {
		return "", terrors.Raise(err, 100012)
	}
	err = as.redis.Set(accessToken, refreshToken)
	if err != nil {
		return "", terrors.Raise(err, 300002)
	}
	ok, err := as.redis.ExpireAt(accessToken, int(duration.Unix()))
	if !ok || err != nil {
		return "", terrors.Raise(err, 300002)
	}
	return accessToken, nil
}

func (as *AuthUC) GenerateTokensPair(params *model.CreateAuthTokensLogic) (accessToken, refreshToken string, err error) {
	refreshDuration := time.Now().Add(time.Minute * time.Duration(as.config.HttpConfig.RefreshExpireTime))
	refreshToken, err = server.CreateJwtToken(&server.JwtParams{
		Salt:     &as.config.HttpConfig.JWTSalt,
		UserId:   &params.UserId,
		Duration: &refreshDuration,
		Type:     ptr.String("refresh"),
	})
	if err != nil {
		return "", "", terrors.Raise(err, 100012)
	}
	accessToken, err = as.GenerateAccessToken(refreshToken, params)
	if err != nil {
		return "", "", err
	}
	err = as.redis.Set(refreshToken, "")
	if err != nil {
		return "", "", terrors.Raise(err, 300002)
	}
	ok, err := as.redis.ExpireAt(refreshToken, int(refreshDuration.Unix()))
	if !ok || err != nil {
		return "", "", terrors.Raise(err, 300002)
	}
	return accessToken, refreshToken, nil
}

func (as *AuthUC) ValidateUser(userId int64) (ok bool, err error) {
	user, err := as.authRepo.FindUserByIdShort(userId)
	if err != nil {
		return false, err
	}
	if user == nil {
		return false, terrors.Raise(err, 100010)
	}

	return true, nil
}

func (as *AuthUC) SignOut(params *model.AuthTokensLogic) error {
	_, err := as.redis.Del(params.AccessToken, params.RefreshToken)
	if err != nil {
		return terrors.Raise(err, 200004)
	}
	return nil
}
