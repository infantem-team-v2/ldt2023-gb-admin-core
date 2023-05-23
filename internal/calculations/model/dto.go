package model

type BaseCalculateRequest struct {
}

type BaseCalculateResponse struct {
	InputData  []*BasicCalculationFieldLogic    `json:"input_data"`
	OutputData []*BasicCategoryCalculationLogic `json:"output_data"`
}
