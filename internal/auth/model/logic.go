package model

type RegistrationDataLogic struct {
	Email            string `json:"email" validate:"required" example:"example@mail.ru"`
	Password         string `json:"password" validate:"required" example:"1234qwerty!"`
	RepeatedPassword string `json:"repeated_password" validate:"required" example:"1234qwerty!"`
}

type PersonalDataLogic struct {
	FullName   *string `json:"full_name" validate:"required" example:"Иванов Иван Иванович"`
	Position   *string `json:"position,omitempty" example:"Старший менеджер по инвестициям"`
	Geographic *struct {
		Country *string `json:"country,omitempty" example:"Российская Федерация"`
		City    *string `json:"city,omitempty" example:"Москва"`
	} `json:"geographic,omitempty"`
}

type BusinessDataLogic struct {
	Name             *string `json:"name,omitempty" example:"ООО ИНФАНТЕМ"`
	Website          *string `json:"website,omitempty" example:"infantem.tech"`
	INN              *string `json:"inn" validate:"required" example:"7707083893"`
	EconomicActivity *string `json:"economic_activity" example:"Производство"`
}

type AuthHeadersLogic struct {
	PublicKey string `json:"public_key"`
	Signature string `json:"signature"`
	Body      []byte `json:"body"`
}

type AuthTokensLogic struct {
	AccessToken  string `json:"-"`
	RefreshToken string `json:"-"`
}

type CreateAuthTokensLogic struct {
	UserId int64 `json:"userId"`
}
