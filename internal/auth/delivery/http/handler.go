package http

import (
	"gb-admin-core/config"
	authInterface "gb-admin-core/internal/auth/interface"
	"gb-admin-core/internal/auth/model"
	"gb-admin-core/internal/pkg/common"
	mdwModel "gb-admin-core/internal/pkg/middleware/model"
	"gb-admin-core/pkg/terrors"
	"gb-admin-core/pkg/thttp/server"
	"github.com/gofiber/fiber/v2"
	"time"
)

type AuthHandler struct {
	AuthUC authInterface.UseCase `di:"authUC"`
	Config *config.Config        `di:"config"`
	prefix string
	router fiber.Router
}

func (ah *AuthHandler) GetRouter() fiber.Router {
	return ah.router
}

func (ah *AuthHandler) GetPrefix() string {
	return ah.prefix
}

func NewAuthHandler(app *fiber.App) server.IHandler {
	return &AuthHandler{
		router: app.Group("auth"),
		prefix: "auth",
	}
}

// setAuthCookies Sets required for authorization cookies w/ tokens
func (ah *AuthHandler) setAuthCookies(ctx *fiber.Ctx, accessToken, refreshToken string) {
	ctx.Cookie(&fiber.Cookie{
		Name:     mdwModel.AccessKey,
		Value:    accessToken,
		Domain:   ah.Config.BaseConfig.Service.URL,
		MaxAge:   int((time.Duration(ah.Config.HttpConfig.AccessExpireTime) * time.Second).Seconds()),
		Secure:   true,
		Path:     "/",
		HTTPOnly: true,
		SameSite: "None",
	})
	ctx.Cookie(&fiber.Cookie{
		Name:     mdwModel.RefreshKey,
		Value:    refreshToken,
		Domain:   ah.Config.BaseConfig.Service.URL,
		MaxAge:   int((time.Duration(ah.Config.HttpConfig.RefreshExpireTime) * time.Minute).Seconds()),
		Secure:   true,
		Path:     "/",
		HTTPOnly: true,
		SameSite: "None",
	})
}

// VendorAuth godoc
// @Summary Sign in or sign up via external vendor
// @Description Accepts token from vendor which we process and returning pair of tokens
// @Tags Authorization
// @Accept json
// @Produce json
// @Param vendor query string true "Vendor which is providing authorization" Enums(apple, google)
// @Success 200 {object} common.Response
// @Failure 400 {object} common.Response
// @Router /auth [post]
func (ah *AuthHandler) VendorAuth() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return terrors.Raise(nil, 300000)
	}
}

// SignUp godoc
// @Summary Sign up with base data
// @Description Sign up with data which was in our task
// @Tags Authorization, Login
// @Accept json
// @Produce json
// @Param data body model.SignUpRequest true "Authorization data from user"
// @Success 201 {object} model.SignUpResponse
// @Failure 400 {object} common.Response
// @Failure 404 {object} common.Response
// @Failure 409 {object} common.Response
// @Failure 422 {object} common.Response
// @Router /auth/sign/up [post]
func (ah *AuthHandler) SignUp() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params model.SignUpRequest
		if err := server.ReadRequest(ctx, &params); err != nil {
			return terrors.Raise(err, 100001)
		}
		response, err := ah.AuthUC.SignUp(&params)
		if err != nil {
			return err
		}
		ah.setAuthCookies(ctx, response.AccessToken, response.RefreshToken)
		return ctx.
			Status(201).
			JSON(response)
	}
}

// SignIn godoc
// @Summary Sign in
// @Description Authorization and get access token
// @Tags Authorization, Login
// @Accept json
// @Produce json
// @Param data body model.SignInRequest true "Authorization data from user"
// @Success 200 {object} model.SignInResponse
// @Failure 400 {object} common.Response
// @Failure 403 {object} common.Response
// @Failure 404 {object} common.Response
// @Failure 422 {object} common.Response
// @Router /auth/sign/in [post]
func (ah *AuthHandler) SignIn() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params model.SignInRequest
		if err := server.ReadRequest(ctx, &params); err != nil {
			return terrors.Raise(err, 100001)
		}
		response, err := ah.AuthUC.SignIn(&params)
		if err != nil {
			return err
		}
		ah.setAuthCookies(ctx, response.AccessToken, response.RefreshToken)
		return ctx.JSON(response)
	}
}

// SignOut godoc
// @Summary Sign out
// @Description Delete tokens
// @Tags Authorization, Login
// @Accept json
// @Produce json
// @Success 200 {object} model.SignOutResponse
// @Failure 400 {object} common.Response
// @Failure 403 {object} common.Response
// @Failure 404 {object} common.Response
// @Failure 422 {object} common.Response
// @Router /auth/sign/out [delete]
func (ah *AuthHandler) SignOut() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		accessToken, refreshToken := ctx.Cookies(mdwModel.AccessKey), ctx.Cookies(mdwModel.RefreshKey)

		ctx.Cookie(&fiber.Cookie{
			Name:     mdwModel.AccessKey,
			Value:    "",
			Domain:   ah.Config.BaseConfig.Service.URL,
			Expires:  time.Now().Add(-time.Hour * 24),
			Secure:   true,
			Path:     "/",
			HTTPOnly: true,
			SameSite: "None",
		})
		ctx.Cookie(&fiber.Cookie{
			Name:     mdwModel.RefreshKey,
			Value:    "",
			Domain:   ah.Config.BaseConfig.Service.URL,
			Expires:  time.Now().Add(-time.Hour * 24),
			Secure:   true,
			Path:     "/",
			HTTPOnly: true,
			SameSite: "None",
		})

		err := ah.AuthUC.SignOut(&model.AuthTokensLogic{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		})
		if err != nil {
			return err
		}
		return ctx.JSON(model.SignOutResponse{
			Response: common.Response{
				InternalCode: 200,
				Message:      "Successful sign out",
			},
		})
	}
}

// ValidateEmail godoc
// @Summary Validating user's email
// @Description Validating user's email with take message on email and writing code
// @Tags Authorization
// @Accept json
// @Produce json
// @Param code body model.EmailValidateRequest true "Data for validation by email from app"
// @Success 200 {object} model.EmailValidateResponse
// @Failure 404 {object} common.Response
// @Router /auth/email/validate [post]
func (ah *AuthHandler) ValidateEmail() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params model.EmailValidateRequest
		if err := server.ReadRequest(ctx, &params); err != nil {
			return terrors.Raise(err, 100001)
		}
		response, err := ah.AuthUC.ValidateEmail(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(response)
	}
}

// ResetPassword godoc
// @Summary Resetting password
// @Description Resetting password by getting validated email for password change
// @Tags Authorization, Password
// @Accept json
// @Produce json
// @Param t-session-key header string true "Session key to identify that this is current session of password change"
// @Param new_pswds body model.ResetPasswordRequest true "New password pair with confirmation to update credentials"
// @Success 200 {object} model.ResetPasswordResponse
// @Failure 400 {object} common.Response
// @Failure 403 {object} common.Response
// @Router /auth/password/reset [put]
func (ah *AuthHandler) ResetPassword() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return nil
	}
}

// PrepareResetPassword godoc
// @Summary Make reset password session
// @Description Creates session for password reset
// @Tags Authorization, Password
// @Accept json
// @Produce json
// @Param new_pswds body model.PrepareResetPasswordRequest true "Session key for backend's session validation"
// @Success 200 {object} model.PrepareResetPasswordResponse
// @Failure 400 {object} common.Response
// @Failure 403 {object} common.Response
// @Router /auth/password/reset/prepare [patch]
func (ah *AuthHandler) PrepareResetPassword() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return nil
	}
}

// ValidateResetPassword godoc
// @Summary Validate reset password session
// @Description Validate reset password session by code that user gets on its email
// @Tags Authorization, Password
// @Accept json
// @Produce json
// @Param t-session-key header string true "Session key to identify that this is current session of password change"
// @Param new_pswds body model.ValidateResetPasswordRequest true "Code from email to validate request"
// @Success 200 {object} model.ValidateResetPasswordResponse
// @Failure 400 {object} common.Response
// @Failure 403 {object} common.Response
// @Router /auth/password/reset/validate [patch]
func (ah *AuthHandler) ValidateResetPassword() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return nil
	}
}

// Check godoc
// @Summary Health and auth check
// @Description Validates that session is authorized
// @Tags Authorization
// @Success 200
// @Failure 400 {object} common.Response
// @Failure 401 {object} common.Response
// @Failure 403 {object} common.Response
// @Router /auth/check [get]
func (ah *AuthHandler) Check() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(200)
	}
}
