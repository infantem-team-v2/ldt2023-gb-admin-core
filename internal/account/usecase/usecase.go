package usecase

import (
	accountInterface "gb-auth-gate/internal/account/interface"
	"gb-auth-gate/internal/account/model"
	"gb-auth-gate/pkg/thttp"
	"github.com/sarulabs/di"
)

type AccountUseCase struct {
	accountRepo accountInterface.RelationalRepository `di:"accountRepo"`
	httpClient  *thttp.ThttpClient                    `di:"httpClient"`
}

func BuildAccountUseCase(ctn di.Container) (interface{}, error) {
	return &AccountUseCase{
		accountRepo: ctn.Get("accountRepo").(accountInterface.RelationalRepository),
		httpClient:  ctn.Get("httpClient").(*thttp.ThttpClient),
	}, nil
}

func (auc *AccountUseCase) GetCommonInfo(userId int64) (*model.GetCommonInfoResponse, error) {
	user, err := auc.accountRepo.GetPersonalUserInfo(userId)
	if err != nil {
		return nil, err
	}
	business, err := auc.accountRepo.GetBusinessInfo(userId)
	if err != nil {
		return nil, err
	}
	return &model.GetCommonInfoResponse{
		PersonalData: model.PersonalDataLogic{
			FullName:    user.FullName,
			Email:       user.Email,
			JobPosition: user.JobPosition,
			Geography:   user.Geography,
		},
		BusinessData: model.BusinessDataLogic{
			Inn:              business.Inn,
			Name:             business.Name,
			EconomicActivity: business.EconomicActivity,
		},
	}, nil
}
