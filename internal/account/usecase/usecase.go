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
	return nil, nil
}
