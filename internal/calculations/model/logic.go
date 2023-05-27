package model

type BasicCompanyLogic struct {
	ProjectName       *string  `json:"project_name"`
	OrganizationType  *string  `json:"organization_type"`
	WorkersQuantity   *int     `json:"workers_quantity"`
	Industry          *string  `json:"industry"`
	County            *string  `json:"county"`
	LandArea          *int     `json:"land_area"`
	BuildingArea      *int     `json:"building_area"`
	MachineNames      []string `json:"machine_names"`
	MachineQuantities []int    `json:"machine_quantities"`
	PatentType        *string  `json:"patent_type"`
	Bookkeeping       *bool    `json:"bookkeeping"`
	TaxSystem         *string  `json:"tax_system"`
	Operations        *int     `json:"operations"`
	OtherNeeds        []string `json:"other_needs"`
}

type MakeCalcRequestLogic struct {
	UserID  *int              `json:"user_id,omitempty"`
	Company BasicCompanyLogic `json:"company"`
}
