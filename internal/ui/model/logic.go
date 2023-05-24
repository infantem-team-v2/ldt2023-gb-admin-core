package model

type UiElementLogic struct {
	Field   string        `json:"field"`
	FieldId string        `json:"field_id"`
	Comment string        `json:"comment"`
	Type    string        `json:"type"`
	Options []interface{} `json:"options"`
}

type UiTypeLogic struct {
	Name            string `json:"type"`
	Comment         string `json:"hint"`
	MultipleOptions bool   `json:"multiple_options"`
}

type UiCategoryLogic struct {
	Category   string            `json:"category"`
	CategoryId string            `json:"category_id"`
	Elements   []*UiElementLogic `json:"elements"`
}
