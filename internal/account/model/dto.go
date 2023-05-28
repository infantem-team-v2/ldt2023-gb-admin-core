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
	Results []struct {
		Name      string `json:"name"`
		Summary   int    `json:"summary"`
		TimeStamp string `json:"time_stamp"`
		ReportID  string `json:"report_id"`
	} `json:"results"`
}

type ChangeConstantsRequest struct {
	Elements []*ChangeConstantUnitLogic `json:"elements"`
}

type GetConstantsResponse struct {
	MachinePrices []struct {
		MachineID    int    `json:"machine_id"`
		MachineName  string `json:"machine_name"`
		MachinePrice int    `json:"machine_price"`
	} `json:"machine_prices"`
	MeanSalaries []struct {
		IndustryID   int    `json:"industry_id"`
		IndustryName string `json:"industry_name"`
		Salary       int    `json:"salary"`
	} `json:"mean_salaries"`
	OtherNeeds []struct {
		NeedID    int     `json:"need_id"`
		NeedName  string  `json:"need_name"`
		NeedCoeff float64 `json:"need_coeff"`
	} `json:"other_needs"`
	PatentPrices []struct {
		PatentID    int    `json:"patent_id"`
		PatentName  string `json:"patent_name"`
		PatentPrice int    `json:"patent_price"`
	} `json:"patent_prices"`
	CountyPrices []struct {
		CountyID    int     `json:"county_id"`
		CountyName  string  `json:"county_name"`
		CountyPrice float64 `json:"county_price"`
	} `json:"county_prices"`
}
