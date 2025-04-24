package handler

import (
	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
	"github.com/NonthapatKim/many-tooth-api/internal/core/domain/response"
	"github.com/gofiber/fiber/v2"
)

func (h *handler) CreateUserInterest(c *fiber.Ctx) error {
	var interestAdd domain.CreateUserInterestRequest

	accessToken := c.Locals("access_token")
	tokenString, ok := accessToken.(string)
	if !ok {
		return response.JSONErrorResponse(c, fiber.StatusUnauthorized, "Invalid or missing access token", nil)
	}

	if err := c.BodyParser(&interestAdd); err != nil {
		return response.JSONErrorResponse(c, fiber.StatusUnauthorized, err.Error(), nil)
	}

	if len(interestAdd.InterestId) == 0 {
		return response.JSONErrorResponse(c, fiber.StatusBadRequest, "No schedule data provided", nil)
	}

	req := domain.CreateUserInterestRequest{
		AccessToken: tokenString,
		InterestId:  interestAdd.InterestId,
	}

	result, err := h.svc.CreateUserInterest(req)
	if err != nil {
		return response.JSONErrorResponse(c, fiber.StatusUnauthorized, err.Error(), nil)
	}

	return response.JSONSuccessResponse(c, result)
}
