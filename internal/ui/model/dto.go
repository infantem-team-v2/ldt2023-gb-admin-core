package model

type GetCalcActiveElementsResponse struct {
	Elements []*UiElementLogic `json:"elements"`
}

type GetTypesResponse struct {
	Elements []*UiTypeLogic `json:"elements"`
}
