package http

import "github.com/gofiber/fiber/v2"

func (mdw *MiddlewareManager) StartLoggingMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {

		return nil
	}
}

func (mdw *MiddlewareManager) EndLoggingMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}
