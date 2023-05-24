package uiInterface

type UseCase interface {
	GetCalcActiveElements() (interface{}, uint16, error)
}
