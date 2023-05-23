package terrors

import (
	"gb-auth-gate/internal/pkg/common"
	"gb-auth-gate/pkg/thttp"
	"gb-auth-gate/pkg/tlogger"
	"github.com/gofiber/fiber/v2"
	"github.com/sarulabs/di"
	"strings"
)

type HttpErrorHandler struct {
	logger tlogger.ILogger `di:"logger"`
}

func BuildErrorHandler(ctn di.Container) (interface{}, error) {
	return &HttpErrorHandler{
		logger: ctn.Get("logger").(tlogger.ILogger),
	}, nil
}

func (heh *HttpErrorHandler) Handle(ctx *fiber.Ctx, err error) error {
	if tErr, ok := err.(*tError); ok {
		requestId := ctx.GetRespHeader(thttp.RequestIdHeader, "no-header-provided")
		heh.logger.ErrorFull(tErr, requestId)
		return ctx.
			Status(tErr.statusCode).
			JSON(tErr.externalMessage)
	}
	if strings.Contains(err.Error(), "Cannot") {
		return ctx.
			Status(404).
			JSON(common.Response{
				InternalCode: 404,
				Message:      err.Error(),
			})
	}
	return ctx.
		Status(500).
		JSON(common.Response{
			InternalCode: 500,
			Message:      err.Error(),
		})
}
