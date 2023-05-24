package http

import (
	uiInterface "gb-auth-gate/internal/ui/interface"
	"gb-auth-gate/pkg/thttp/server"
	"github.com/gofiber/fiber/v2"
)

type UiHandler struct {
	UiUC   uiInterface.UseCase `di:"uiUC"`
	prefix string
	router fiber.Router
}

func NewUiHandler(app *fiber.App) server.IHandler {
	return &UiHandler{
		prefix: "ui",
		router: app.Group("ui"),
	}
}

func (uh *UiHandler) GetRouter() fiber.Router {
	return uh.router
}

func (uh *UiHandler) GetPrefix() string {
	return uh.prefix
}

// GetCalcActiveElements godoc
// @Summary Get UI elements for calculator
// @Description Get specification for ui elements to visualise it on frontend
// @Tags UI
// @Success 200 {object} model.GetCalcActiveElementsResponse
// @Failure 400 {object} common.Response
// @Failure 401 {object} common.Response
// @Failure 403 {object} common.Response
// @Failure 422 {object} common.Response
// @Failure 409 {object} common.Response
// @Router /ui/calc/element/active [get]
func (uh *UiHandler) GetCalcActiveElements() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		response, statusCode, err := uh.UiUC.GetCalcActiveElements()
		if err != nil {
			return err
		}
		return ctx.Status(int(statusCode)).JSON(response)
	}
}
