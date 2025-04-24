package middleware

import (
	"strings"

	"github.com/NonthapatKim/many-tooth-api/internal/core/domain/response"
	"github.com/gofiber/fiber/v2"
)

func Authorization() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if authHeader != "" {
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				return response.JSONErrorResponse(c, fiber.StatusUnauthorized, "invalid authorization header format", nil)
			}

			tokenString := parts[1]
			c.Locals("access_token", tokenString)
		}

		return c.Next()
	}
}
