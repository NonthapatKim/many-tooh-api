package response

import (
	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
	"github.com/gofiber/fiber/v2"
)

func JSONErrorResponse(c *fiber.Ctx, statusCode int, message string, errors *domain.ValidationErrorResponse) error {
	response := fiber.Map{
		"code": statusCode,
	}

	if message != "" {
		response["message"] = message
	}

	if errors != nil {
		response["errors"] = *errors
	}

	return c.Status(statusCode).JSON(response)
}

func JSONSuccessResponse(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(data)
}
