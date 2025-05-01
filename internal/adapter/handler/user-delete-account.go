package handler

import (
	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
	"github.com/NonthapatKim/many-tooth-api/internal/core/domain/response"
	"github.com/gofiber/fiber/v2"
)

func (h *handler) UserDeleteAccount(c *fiber.Ctx) error {
	var userDelete domain.UserDeleteAccountRequest

	accessToken := c.Locals("access_token")
	tokenString, ok := accessToken.(string)
	if !ok {
		return response.JSONErrorResponse(c, fiber.StatusUnauthorized, "Invalid or missing access token", nil)
	}

	if err := c.BodyParser(&userDelete); err != nil {
		return response.JSONErrorResponse(c, fiber.StatusUnauthorized, err.Error(), nil)
	}

	req := domain.UserDeleteAccountRequest{
		LocalDeviceToken: userDelete.LocalDeviceToken,
		AccessToken:      tokenString,
	}

	result, err := h.svc.UserDeleteAccount(req)
	if err != nil {
		return response.JSONErrorResponse(c, fiber.StatusUnauthorized, err.Error(), nil)
	}

	return response.JSONSuccessResponse(c, result)
}
