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
	return &model.BaseCalculateResponse{
		InputData: []*model.BasicCalculationFieldLogic{
			{
				Field: "Отрасль",
				Value: "Фармацевтика",
			},
		},
		OutputData: []*model.BasicCategoryCalculationLogic{
			{
				Category: "Персонал",
				Data: []*model.BasicCalculationFieldLogic{
					{
						Field: "Затраты на зарплату",
						Value: "100 тыс. руб.",
					},
				},
			},
		},
	}, nil
}
