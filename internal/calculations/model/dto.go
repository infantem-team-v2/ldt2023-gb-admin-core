package model

type BaseCalculateRequest struct {
	BasicCompanyLogic
}

type ImprovedCalculateResponse struct {
	BaseCalculateResponse
	Link string `json:"link"`
}

type GetCalculatorConstantResponse struct {
	Data struct {
		Machines   []string `json:"machines"`
		Industries []string `json:"industries"`
		Needs      []string `json:"needs"`
		Patents    []string `json:"patents"`
	} `json:"data"`
}

type BaseCalculateResponse struct {
	TrackerID     string `json:"tracker_id" rus:"tracker_id"`
	TotalExpenses int    `json:"total_expenses" rus:"total_expenses"`
	Output        struct {
		Service struct {
			ServiceExpenses     int `json:"service_expenses" rus:"Суммарные_затраты_на_услуги"`
			DutyExpenses        int `json:"duty_expenses" rus:"Пошлина"`
			BookkeepingExpenses int `json:"bookkeeping_expenses" rus:"Бухгалтерия"`
			PatentExpenses      int `json:"patent_expenses" rus:"Патент"`
			MachineExpenses     int `json:"machine_expenses" rus:"Оборудование"`
		} `json:"service" rus:"Услуги"`
		Estate struct {
			EstateExpenses   int `json:"estate_expenses" rus:"Суммарные_затраты_на_недвижимость"`
			LandExpenses     int `json:"land_expenses" rus:"Затраты_на_землю"`
			BuildingExpenses int `json:"building_expenses" rus:"Затраты_на_строительство"`
		} `json:"estate" rus:"Недвижимость"`
		Staff struct {
			StaffExpenses    int `json:"staff_expenses" rus:"Суммарные_затраты_на_персонал"`
			SalariesExpenses int `json:"salaries_expenses" rus:"Зарплаты"`
			PensionExpenses  int `json:"pension_expenses" rus:"Пенсионные_отчисления"`
			MedicalExpenses  int `json:"medical_expenses" rus:"Медицинские_отчисления"`
		} `json:"staff" rus:"Персонал"`
		Tax struct {
			TaxExpenses int `json:"tax_expenses" rus:"Суммарные_затраты_на_налоги_за_год"`
			LandTax     int `json:"land_tax" rus:"Налог_на_землю"`
			EstateTax   int `json:"estate_tax" rus:"Налог_на_недвижимость"`
			IncomeTax   int `json:"income_tax" rus:"Налог_на_доход"`
		} `json:"tax" rus:"Налоги"`
	} `json:"output" rus:"output"`
	Input struct {
		ProjectName       *string  `json:"project_name,omitempty" rus:"Название_компании"`
		OrganizationType  *string  `json:"organization_type,omitempty" rus:"Тип_организации"`
		WorkersQuantity   *int     `json:"workers_quantity,omitempty" rus:"Колиество_сотрудников"`
		Industry          *string  `json:"industry,omitempty" rus:"Отрасль_производства"`
		County            *string  `json:"county,omitempty" rus:"Административный_округ"`
		LandArea          *int     `json:"land_area,omitempty" rus:"Площадь_участка"`
		BuildingArea      *int     `json:"building_area,omitempty" rus:"Площадь_строительства"`
		MachineNames      []string `json:"machine_names,omitempty" rus:"Оборудование"`
		MachineQuantities []int    `json:"machine_quantities,omitempty" rus:"Количество_оборудования"`
		PatentType        *string  `json:"patent_type,omitempty" rus:"Патент"`
		Bookkeeping       *bool    `json:"bookkeeping,omitempty" rus:"Бухгалтерия"`
		TaxSystem         *string  `json:"tax_system,omitempty" rus:"Система_налогообложения"`
		Operations        *int     `json:"operations,omitempty" rus:"Количество_бухгалтерских_операций"`
		OtherNeeds        []string `json:"other_needs,omitempty" rus:"Иные_потребности"`
	} `json:"input" rus:"input"`
}

type GetInsightsResponse struct {
	UsualExpensesInsight struct {
		Insight string `json:"insight"`
	} `json:"usual_expenses_insight"`
	UsualCountyInsight struct {
		Insight string `json:"insight"`
	} `json:"usual_county_insight"`
	WorkersQuantityInsight struct {
		Insight string `json:"insight"`
	} `json:"workers_quantity_insight"`
	BestTaxSystemInsight struct {
		Insight string `json:"insight"`
	} `json:"best_tax_system_insight"`
}

type GetPlotsResponse struct {
	ExpensesDistribution PieChartLogic `json:"expenses_distribution"`
	TaxesDistribution    PieChartLogic `json:"taxes_distribution"`
	PopularityChart      struct {
		Labels   []string           `json:"labels"`
		Datasets []map[string][]int `json:"datasets"`
	} `json:"popularity_chart"`
}
