package model

type BasicCompanyLogic struct {
	ProjectName       *string  `json:"project_name,omitempty"`
	OrganizationType  *string  `json:"organization_type,omitempty"`
	WorkersQuantity   *int     `json:"workers_quantity,omitempty"`
	Industry          *string  `json:"industry,omitempty"`
	County            *string  `json:"county,omitempty"`
	LandArea          *int     `json:"land_area,omitempty"`
	BuildingArea      *int     `json:"building_area,omitempty"`
	MachineNames      []string `json:"machine_names,omitempty"`
	MachineQuantities []int    `json:"machine_quantities,omitempty"`
	PatentType        *string  `json:"patent_type,omitempty"`
	Bookkeeping       *bool    `json:"bookkeeping,omitempty"`
	TaxSystem         *string  `json:"tax_system,omitempty"`
	Operations        *int     `json:"operations,omitempty"`
	OtherNeeds        []string `json:"other_needs,omitempty"`
}

type MakeCalcRequestLogic struct {
	UserID  *int              `json:"user_id,omitempty"`
	Company BasicCompanyLogic `json:"company"`
}
