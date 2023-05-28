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

type CalcResponseLogic struct {
	TrackerID     string `json:"tracker_id"`
	TotalExpenses int    `json:"total_expenses"`
	Output        struct {
		Service struct {
			ServiceExpenses     int `json:"service_expenses"`
			DutyExpenses        int `json:"duty_expenses"`
			BookkeepingExpenses int `json:"bookkeeping_expenses"`
			PatentExpenses      int `json:"patent_expenses"`
			MachineExpenses     int `json:"machine_expenses"`
		} `json:"service"`
		Estate struct {
			EstateExpenses   int `json:"estate_expenses"`
			LandExpenses     int `json:"land_expenses"`
			BuildingExpenses int `json:"building_expenses"`
		} `json:"estate"`
		Staff struct {
			StaffExpenses    int `json:"staff_expenses"`
			SalariesExpenses int `json:"salaries_expenses"`
			PensionExpenses  int `json:"pension_expenses"`
			MedicalExpenses  int `json:"medical_expenses"`
		} `json:"staff"`
		Tax struct {
			TaxExpenses int `json:"tax_expenses"`
			LandTax     int `json:"land_tax"`
			EstateTax   int `json:"estate_tax"`
			IncomeTax   int `json:"income_tax"`
		} `json:"tax"`
	} `json:"output"`
	Input struct {
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
	} `json:"input"`
}

type PieChartLogic struct {
	Labels   []string         `json:"labels"`
	Datasets []map[string]int `json:"datasets"`
}
