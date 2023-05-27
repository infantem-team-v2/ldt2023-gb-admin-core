package http

import (
	calculationsInterface "gb-auth-gate/internal/calculations/interface"
	"gb-auth-gate/internal/calculations/model"
	mdwModel "gb-auth-gate/internal/pkg/middleware/model"
	"gb-auth-gate/pkg/terrors"
	"gb-auth-gate/pkg/thttp/server"
	"gb-auth-gate/pkg/tutils/ptr"
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
		response, err := ch.CalculationsUC.BaseCalculate(&params, nil)
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
		userId := ctx.Locals(mdwModel.UserIdLocals).(int64)
		response, err := ch.CalculationsUC.ImprovedCalculate(&params, ptr.Int(int(userId)))

		if err != nil {
			return err
		}

		return ctx.JSON(response)
	}
}

// GetResult godoc
// @Summary Get result by tracker id
// @Description Get report by tracker id
// @Tags Calculator
// @Success 200 {object} model.BaseCalculateRequest
// @Failure 400 {object} common.Response
// @Failure 401 {object} common.Response
// @Failure 403 {object} common.Response
// @Failure 404 {object} common.Response
// @Failure 409 {object} common.Response
// @Failure 422 {object} common.Response
// @Router /calc/report/:trackerId [get]
func (ch *CalculationsHandler) GetResult() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		trackerId := ctx.Params("trackerId", "")
		if trackerId == "" {
			return terrors.Raise(nil, 100019)
		}
		response, err := ch.CalculationsUC.GetResult(trackerId)
		if err != nil {
			return err
		}
		return ctx.JSON(response)
	}
}

// GetCalculatorConstant godoc
// @Summary Get constants for calculator's fields
// @Description Get report by tracker id
// @Tags Calculator
// @Success 200 {object} model.GetCalculatorConstantResponse
// @Failure 400 {object} common.Response
// @Failure 401 {object} common.Response
// @Failure 403 {object} common.Response
// @Failure 404 {object} common.Response
// @Failure 409 {object} common.Response
// @Failure 422 {object} common.Response
// @Router /calc/fields [get]
func (ch *CalculationsHandler) GetCalculatorConstant() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		response, err := ch.CalculationsUC.GetConstants()
		if err != nil {
			return err
		}
		return ctx.JSON(response)
	}
}
