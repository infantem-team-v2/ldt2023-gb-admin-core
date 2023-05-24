package model

type BaseCalculateRequest struct {
	Inputs []*BasicCalculationFieldLogic `json:"inputs"`
}

type BaseCalculateResponse struct {
	InputData  []*BasicCalculationFieldLogic    `json:"input_data"`
	OutputData []*BasicCategoryCalculationLogic `json:"output_data"`
}

type ImprovedCalculateResponse struct {
	BaseCalculateResponse
	Link string `json:"link"`
}
