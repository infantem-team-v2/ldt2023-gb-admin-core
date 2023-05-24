package model

type PersonalDataLogic struct {
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	JobPosition string `json:"job_position"`
	Geography   string `json:"geography"`
}

type BusinessDataLogic struct {
	Inn              string `json:"inn"`
	Name             string `json:"name"`
	EconomicActivity string `json:"economic_activity"`
	Website          string `json:"website"`
}
