package http

import "github.com/gofiber/fiber/v2"

type AccountHandler struct {
	prefix string
	router fiber.Router
}

func (ah *AccountHandler) GetRouter() fiber.Router {
	return ah.router
}

func (ah *AccountHandler) GetPrefix() string {
	return ah.prefix
}
