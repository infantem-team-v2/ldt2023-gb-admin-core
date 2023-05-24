package model

type GetCalcActiveElementsResponse struct {
	Elements []*UiCategoryLogic `json:"categories"`
}

type GetTypesResponse struct {
	Elements []*UiTypeLogic `json:"elements"`
}
