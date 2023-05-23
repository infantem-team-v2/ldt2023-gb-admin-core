package model

type BasicCalculationFieldLogic struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

type BasicCategoryCalculationLogic struct {
	Category string                        `json:"category"`
	Data     []*BasicCalculationFieldLogic `json:"data"`
}
