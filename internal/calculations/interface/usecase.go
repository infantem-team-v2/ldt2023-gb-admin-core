package calculationsInterface

import "gb-auth-gate/internal/calculations/model"

type UseCase interface {
	BaseCalculate(params *model.BaseCalculateRequest) (*model.BaseCalculateResponse, error)
}
