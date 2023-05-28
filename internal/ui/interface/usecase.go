package uiInterface

import "gb-admin-core/internal/ui/model"

type UseCase interface {
	GetCalcActiveElements() (interface{}, uint16, error)
	SetActiveForElements(params *model.SetActiveForElementRequest) (interface{}, uint16, error)
}
