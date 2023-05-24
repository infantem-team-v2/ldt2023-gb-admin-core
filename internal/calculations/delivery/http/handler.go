package http

import (
	calculationsInterface "gb-auth-gate/internal/calculations/interface"
	"gb-auth-gate/internal/calculations/model"
	"gb-auth-gate/pkg/terrors"
	"gb-auth-gate/pkg/thttp/server"
	"github.com/gofiber/fiber/v2"
)

type CalculationsHandler struct {
	CalculationsUC calculationsInterface.UseCase `di:"calcUC"`
	prefix         string
	router         fiber.Router
}

func NewCalculationsHandler(app *fiber.App) server.IHandler {
	return &CalculationsHandler{
		router: app.Group("calc"),
		prefix: "calc",
	}
}

func (ch *CalculationsHandler) GetRouter() fiber.Router {
	return ch.router
}

func (ch *CalculationsHandler) GetPrefix() string {
	return ch.prefix
}

// BaseCalculate godoc
// @Summary Base calculation
// @Description Base calculation without authorization for landing page
// @Tags Calculator
// @Accept json
// @Produce json
// @Param data body model.BaseCalculateRequest true "Basic parameters for base calculator w/o auth"
// @Success 200 {object} model.BaseCalculateResponse
// @Failure 400 {object} common.Response
// @Router /calc/base [post]
func (ch *CalculationsHandler) BaseCalculate() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params model.BaseCalculateRequest
		if err := server.ReadRequest(ctx, &params); err != nil {
			return terrors.Raise(err, 100001)
		}
		response, err := ch.CalculationsUC.BaseCalculate(&params)
		if err != nil {
			return err
		}

		return ctx.JSON(response)
	}
}

// ImprovedCalculate godoc
// @Summary Improved calculation w/ auth
// @Description Calculations with authorization
// @Tags Calculator
// @Accept json
// @Produce json
// @Param data body model.BaseCalculateRequest true "Basic parameters for base calculator w/o auth"
// @Success 200 {object} model.ImprovedCalculateResponse
// @Failure 400 {object} common.Response
// @Router /calc [post]
func (ch *CalculationsHandler) ImprovedCalculate() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params model.BaseCalculateRequest
		if err := server.ReadRequest(ctx, &params); err != nil {
			return terrors.Raise(err, 100001)
		}

		response, err := ch.CalculationsUC.BaseCalculate(&params)
		if err != nil {
			return err
		}

		return ctx.JSON(response)
	}
}
