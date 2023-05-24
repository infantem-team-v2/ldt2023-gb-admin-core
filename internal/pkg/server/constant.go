package server

import (
	accountHttp "gb-auth-gate/internal/account/delivery/http"
	authHttp "gb-auth-gate/internal/auth/delivery/http"
	calcHttp "gb-auth-gate/internal/calculations/delivery/http"
	uiHttp "gb-auth-gate/internal/ui/delivery/http"
	"gb-auth-gate/pkg/thttp/server"
	"github.com/gofiber/fiber/v2"
)

var (
	emptyHandlers = map[string]func(app *fiber.App) server.IHandler{
		"account": accountHttp.NewAccountHandler,
		"auth":    authHttp.NewAuthHandler,
		"calc":    calcHttp.NewCalculationsHandler,
		"ui":      uiHttp.NewUiHandler,
	}
)
