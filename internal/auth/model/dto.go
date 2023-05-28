package model

import (
	"gb-admin-core/internal/pkg/common"
)

type SignUpRequest struct {
	//server.Params
	AuthTokensLogic
	AuthData     RegistrationDataLogic `json:"auth_data"`
	PersonalData PersonalDataLogic     `json:"personal_data"`
	BusinessData BusinessDataLogic     `json:"business_data"`
}

type SignUpResponse struct {
	AuthTokensLogic
	common.Response
}

type EmailValidateRequest struct {
	Code int32 `json:"code" validate:"required"`
}

type EmailValidateResponse struct {
	common.Response
	Valid bool `json:"valid"`
}

type SignInRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password"  validate:"required"`
}

type SignInResponse struct {
	common.Response
	AuthTokensLogic
	Email string `json:"email"`
}

type SignOutResponse struct {
	common.Response
}

type ResetPasswordRequest struct {
	Password         string `json:"password" validate:"required"`
	RepeatedPassword string `json:"repeated_password" validate:"required"`
}

type ResetPasswordResponse struct {
	common.Response
}

type PrepareResetPasswordRequest struct {
	Email string `json:"email" validate:"required"`
}

type PrepareResetPasswordResponse struct {
	common.Response
	SessionKey string `json:"session_key"`
}

type ValidateResetPasswordRequest struct {
	ValidationCode string `json:"validation_code" validate:"required"`
}

type ValidateResetPasswordResponse struct {
	common.Response
}
