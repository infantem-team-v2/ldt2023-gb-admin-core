package model

import "time"

type GetCommonInfoResponse struct {
	PersonalData PersonalDataLogic `json:"personal_data"`
	BusinessData BusinessDataLogic `json:"business_data"`
}

type UpdateUserInfoRequest struct {
	PersonalData PersonalDataLogic `json:"personal_data"`
	BusinessData BusinessDataLogic `json:"business_data"`
}

type GetResultsByUserResponse struct {
	Results []struct {
		Name      string    `json:"name"`
		Summary   int       `json:"summary"`
		TimeStamp time.Time `json:"time_stamp"`
		ReportID  string    `json:"report_id"`
	} `json:"results"`
}
