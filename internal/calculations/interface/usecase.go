package calculationsInterface

import "gb-auth-gate/internal/calculations/model"

type UseCase interface {
	BaseCalculate(params *model.BaseCalculateRequest, userId *int) (*model.BaseCalculateResponse, error)
	ImprovedCalculate(params *model.BaseCalculateRequest, userId *int) (*model.ImprovedCalculateResponse, error)
}
