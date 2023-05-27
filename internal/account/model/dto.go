package model

type GetCommonInfoResponse struct {
	PersonalData PersonalDataLogic `json:"personal_data"`
	BusinessData BusinessDataLogic `json:"business_data"`
}

type UpdateUserInfoRequest struct {
	PersonalData PersonalDataLogic `json:"personal_data"`
	BusinessData BusinessDataLogic `json:"business_data"`
}

type GetResultsByUserResponse struct {
}
