package server

import (
	authHttp "gb-auth-gate/internal/auth/delivery/http"
	calcHttp "gb-auth-gate/internal/calculations/delivery/http"
	"gb-auth-gate/pkg/thttp/server"
	"github.com/gofiber/fiber/v2"
)

var (
	emptyHandlers = map[string]func(app *fiber.App) server.IHandler{
		"auth": authHttp.NewAuthHandler,
		"calc": calcHttp.NewCalculationsHandler,
	}
)
