package usecase

import (
	"gb-auth-gate/internal/calculations/model"
	"gb-auth-gate/pkg/thttp"
	"github.com/sarulabs/di"
)

type CalculationsUseCase struct {
	httpClient *thttp.ThttpClient `di:"httpClient"`
}

func BuildCalculationsUseCase(ctn di.Container) (interface{}, error) {
	return &CalculationsUseCase{
		httpClient: ctn.Get("httpClient").(*thttp.ThttpClient),
	}, nil
}

func (cu *CalculationsUseCase) BaseCalculate(params *model.BaseCalculateRequest) (*model.BaseCalculateResponse, error) {

	return &model.BaseCalculateResponse{}, nil

}

func (cu *CalculationsUseCase) ImprovedCalculate(params *model.BaseCalculateRequest) (*model.ImprovedCalculateResponse, error) {
	baseResponse, err := cu.BaseCalculate(params)
	if err != nil {
		return nil, err
	}
	return &model.ImprovedCalculateResponse{
		BaseCalculateResponse: *baseResponse,
		Link:                  "http://abobus.amogus/trahni_psa.pdf",
	}, nil
}
