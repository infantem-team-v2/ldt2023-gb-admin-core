package model

type BasicCalculationFieldLogic struct {
	Field string      `json:"field"`
	Value interface{} `json:"value"`
}

type BasicCategoryCalculationLogic struct {
	Category string                        `json:"category"`
	Data     []*BasicCalculationFieldLogic `json:"data"`
}
