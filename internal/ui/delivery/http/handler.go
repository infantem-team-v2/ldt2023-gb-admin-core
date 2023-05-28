package http

import (
	uiInterface "gb-admin-core/internal/ui/interface"
	"gb-admin-core/internal/ui/model"
	"gb-admin-core/pkg/terrors"
	"gb-admin-core/pkg/thttp/server"
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

// SetActiveForElement godoc
// @Summary Set active/inactive state for element
// @Description Set state of activity for element
// @Tags UI, Admin
// @Param data body model.SetActiveForElementRequest true "Fields and their states"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.Response
// @Failure 401 {object} common.Response
// @Failure 403 {object} common.Response
// @Failure 422 {object} common.Response
// @Failure 409 {object} common.Response
// @Router /ui/calc/element/active [patch]
func (uh *UiHandler) SetActiveForElement() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params model.SetActiveForElementRequest
		if err := server.ReadRequest(ctx, &params); err != nil {
			return terrors.Raise(err, 100001)
		}
		response, statusCode, err := uh.UiUC.SetActiveForElements(&params)
		if err != nil {
			return err
		}
		return ctx.Status(int(statusCode)).JSON(response)

	}
}
