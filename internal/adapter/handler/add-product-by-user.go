package handler

import (
	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
	"github.com/NonthapatKim/many-tooth-api/internal/core/domain/response"
	"github.com/gofiber/fiber/v2"
)

func (h *handler) AddProductByUser(c *fiber.Ctx) error {
	accessToken := c.Locals("access_token")
	tokenString, ok := accessToken.(string)
	if !ok {
		return response.JSONErrorResponse(c, fiber.StatusUnauthorized, "Invalid or missing access token", nil)
	}

	file, err := c.FormFile("product_image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to get image file",
		})
	}

	src, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to open image file",
		})
	}
	defer src.Close()

	req := domain.AddProductByUserRequest{
		AccessToken: tokenString,
		ProductName: c.FormValue("product_name"),
		Image:       src,
	}

	result, err := h.svc.AddProductByUser(req)
	if err != nil {
		if validationErrs, ok := err.(domain.ValidationError); ok {
			return response.JSONErrorResponse(c, fiber.StatusBadRequest, "", &validationErrs.Errors)
		}

		return response.JSONErrorResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return response.JSONSuccessResponse(c, result)
}
