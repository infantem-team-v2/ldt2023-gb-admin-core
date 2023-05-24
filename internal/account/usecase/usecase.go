package usecase

import (
	"context"
	accountInterface "gb-auth-gate/internal/account/interface"
	"gb-auth-gate/internal/account/model"
	"gb-auth-gate/pkg/thttp"
	"github.com/sarulabs/di"
	"strings"
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
			Website:          business.Website,
		},
	}, nil
}

func (auc *AccountUseCase) UpdateUserInfo(userId int64, params *model.UpdateUserInfoRequest) (*model.GetCommonInfoResponse, error) {
	geoSlice := strings.Split(params.PersonalData.Geography, ", ")
	if len(geoSlice) < 2 {
		if len(geoSlice) < 1 {
			geoSlice = []string{"Москва", "Россия"}
		} else {
			geoSlice = append(geoSlice, "Россия")
		}
	}
	err := auc.accountRepo.UpdatePersonalInfo(context.Background(), &model.UpdateUserDataDAO{
		UserId: userId,

		FullName:    params.PersonalData.FullName,
		Email:       params.PersonalData.Email,
		JobPosition: params.PersonalData.JobPosition,

		City:    geoSlice[0],
		Country: geoSlice[1],

		Inn:              params.BusinessData.Inn,
		Name:             params.BusinessData.Name,
		EconomicActivity: params.BusinessData.EconomicActivity,
		Website:          params.BusinessData.Website,
	})
	if err != nil {
		return nil, err
	}

	return auc.GetCommonInfo(userId)
}
