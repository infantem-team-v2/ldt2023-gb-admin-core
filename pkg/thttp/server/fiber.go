package server

import (
	"gb-auth-gate/config"
	"gb-auth-gate/pkg/terrors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sarulabs/di"
	"time"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func BuildFiberApp(ctn di.Container) (interface{}, error) {
	cfg := ctn.Get("config").(*config.Config)
	errorHandler := ctn.Get("errorHandler").(*terrors.HttpErrorHandler)
	return fiber.New(fiber.Config{
		AppName:               cfg.BaseConfig.Service.Name,
		DisableStartupMessage: true,
		//Prefork:               true,
		ErrorHandler: errorHandler.Handle,
		WriteTimeout: time.Duration(cfg.HttpConfig.TimeOut) * time.Second,
		ReadTimeout:  time.Duration(cfg.HttpConfig.TimeOut) * time.Second,
	}), nil
}

func ReadRequest(c *fiber.Ctx, request interface{}) error {
	if err := c.BodyParser(request); err != nil {
		return err
	}

	return validate.StructCtx(c.Context(), request)
}
