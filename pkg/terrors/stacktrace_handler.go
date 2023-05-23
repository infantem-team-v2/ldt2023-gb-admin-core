package terrors

import (
	"gb-auth-gate/pkg/tlogger"
	"github.com/gofiber/fiber/v2"
	"github.com/sarulabs/di"
)

type StacktraceHandler struct {
	logger       tlogger.ILogger
	errorHandler *HttpErrorHandler
}

func BuildStacktraceHandler(ctn di.Container) (interface{}, error) {
	return &StacktraceHandler{
		logger:       ctn.Get("logger").(tlogger.ILogger),
		errorHandler: ctn.Get("errorHandler").(*HttpErrorHandler),
	}, nil
}

func (sh *StacktraceHandler) Handle(c *fiber.Ctx, e interface{}) {
	if err, ok := e.(error); ok {
		sh.errorHandler.Handle(c, err)
	} else {
		sh.logger.Errorf("Panic: %v", e)
	}

	return
}
