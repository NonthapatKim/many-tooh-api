package middleware

import (
	"github.com/NonthapatKim/many-tooth-api/infrastructure/logs"
	"github.com/gofiber/fiber/v2"
)

func LoggerMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		method := c.Method()
		url := c.OriginalURL()

		logs.LogMiddleWareInfo("Request: %s %s", method, url)
		return c.Next()
	}
}
