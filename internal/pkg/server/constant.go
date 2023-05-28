package server

import (
	accountHttp "gb-admin-core/internal/account/delivery/http"
	authHttp "gb-admin-core/internal/auth/delivery/http"
	uiHttp "gb-admin-core/internal/ui/delivery/http"
	"gb-admin-core/pkg/thttp/server"
	"github.com/gofiber/fiber/v2"
)

var (
	emptyHandlers = map[string]func(app *fiber.App) server.IHandler{
		"account": accountHttp.NewAccountHandler,
		"auth":    authHttp.NewAuthHandler,
		"ui":      uiHttp.NewUiHandler,
	}
)
