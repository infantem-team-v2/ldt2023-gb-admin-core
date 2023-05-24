package http

import (
	accountInterface "gb-auth-gate/internal/account/interface"
	mdwModel "gb-auth-gate/internal/pkg/middleware/model"
	"gb-auth-gate/pkg/thttp/server"
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

func (ah *AccountHandler) UpdateCommonInfo() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return nil
	}
}
