package http

import (
	accountInterface "gb-admin-core/internal/account/interface"
	"gb-admin-core/internal/account/model"
	mdwModel "gb-admin-core/internal/pkg/middleware/model"
	"gb-admin-core/pkg/terrors"
	"gb-admin-core/pkg/thttp/server"
	"github.com/gofiber/fiber/v2"
)

type AccountHandler struct {
	AccountUC accountInterface.UseCase `di:"accountUC"`
	prefix    string
	router    fiber.Router
}

func NewAccountHandler(app *fiber.App) server.IHandler {
	return &AccountHandler{
		prefix: "account",
		router: app.Group("account"),
	}
}

func (ah *AccountHandler) GetRouter() fiber.Router {
	return ah.router
}

func (ah *AccountHandler) GetPrefix() string {
	return ah.prefix
}

// GetCommonInfo godoc
// @Summary Get information about user
// @Description Endpoint to get information about user
// @Tags Account
// @Success 200 {object} model.GetCommonInfoResponse
// @Failure 400 {object} common.Response
// @Failure 401 {object} common.Response
// @Failure 403 {object} common.Response
// @Failure 404 {object} common.Response
// @Failure 409 {object} common.Response
// @Router /account/info [get]
func (ah *AccountHandler) GetCommonInfo() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userId := ctx.Locals(mdwModel.UserIdLocals).(int64)
		response, err := ah.AccountUC.GetCommonInfo(userId)
		if err != nil {
			return err
		}
		return ctx.JSON(response)
	}
}

// UpdateCommonInfo godoc
// @Summary Update user's information
// @Description Endpoint to get information about user
// @Tags Account
// @Param updated_data body model.UpdateUserInfoRequest true "Updated information about user"
// @Success 200 {object} model.GetCommonInfoResponse
// @Failure 400 {object} common.Response
// @Failure 401 {object} common.Response
// @Failure 403 {object} common.Response
// @Failure 404 {object} common.Response
// @Failure 409 {object} common.Response
// @Router /account/info [put]
func (ah *AccountHandler) UpdateCommonInfo() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userId := ctx.Locals(mdwModel.UserIdLocals).(int64)
		var params model.UpdateUserInfoRequest
		if err := server.ReadRequest(ctx, &params); err != nil {
			return terrors.Raise(err, 100001)
		}
		response, err := ah.AccountUC.UpdateUserInfo(userId, &params)
		if err != nil {
			return err
		}
		return ctx.JSON(response)
	}
}

// GetResultsForAccount godoc
// @Summary Get all results for account
// @Description Endpoint to get all information by user (temp w/o pagination)
// @Tags Account
// @Success 200 {object} model.GetResultsByUserResponse
// @Failure 400 {object} common.Response
// @Failure 401 {object} common.Response
// @Failure 403 {object} common.Response
// @Failure 404 {object} common.Response
// @Failure 409 {object} common.Response
// @Failure 422 {object} common.Response
// @Router /account/results [get]
func (ah *AccountHandler) GetResultsForAccount() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userId := ctx.Locals(mdwModel.UserIdLocals).(int64)
		response, statusCode, err := ah.AccountUC.GetResultsByUser(userId)
		if err != nil {
			return err
		}

		return ctx.Status(int(statusCode)).JSON(response)
	}
}

// ChangeConstants godoc
// @Summary Change constants for formula
// @Description Endpoint to change constants
// @Tags Account, Admin
// @Param updated_data body model.ChangeConstantsRequest true "Change constants body"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.Response
// @Failure 401 {object} common.Response
// @Failure 403 {object} common.Response
// @Failure 404 {object} common.Response
// @Failure 409 {object} common.Response
// @Router /account/constant [patch]
func (ah *AccountHandler) ChangeConstants() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params model.ChangeConstantsRequest
		if err := server.ReadRequest(ctx, &params); err != nil {
			return terrors.Raise(err, 100001)
		}
		response, statusCode, err := ah.AccountUC.UpdateConstants(&params)
		if err != nil {
			return err
		}
		return ctx.Status(int(statusCode)).JSON(response)
	}
}

// GetConstants godoc
// @Summary Get all constants for formula
// @Description
// @Tags Account, Admin
// @Success 200 {object} model.GetConstantsResponse
// @Failure 400 {object} common.Response
// @Failure 401 {object} common.Response
// @Failure 403 {object} common.Response
// @Failure 404 {object} common.Response
// @Failure 409 {object} common.Response
// @Failure 422 {object} common.Response
// @Router /account/constant [get]
func (ah *AccountHandler) GetConstants() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		response, statusCode, err := ah.AccountUC.GetConstants()

		if err != nil {
			return err
		}

		return ctx.Status(int(statusCode)).JSON(response)
	}
}
