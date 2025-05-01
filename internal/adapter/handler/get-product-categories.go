package handler

import (
	"github.com/NonthapatKim/many-tooth-api/internal/core/domain/response"
	"github.com/gofiber/fiber/v2"
)

func (h *handler) GetProductCategories(c *fiber.Ctx) error {
	result, err := h.svc.GetProductCategories()
	if err != nil {
		return response.JSONErrorResponse(c, fiber.StatusUnauthorized, err.Error(), nil)
	}

	return response.JSONSuccessResponse(c, result)
}
