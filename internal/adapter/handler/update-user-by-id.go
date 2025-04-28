package handler

import (
	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
	"github.com/NonthapatKim/many-tooth-api/internal/core/domain/response"
	"github.com/gofiber/fiber/v2"
)

func (h *handler) UpdateUserById(c *fiber.Ctx) error {
	var updateUserReq domain.UpdateUserByIdRequest

	accessToken := c.Locals("access_token")
	tokenString, ok := accessToken.(string)
	if !ok {
		return response.JSONErrorResponse(c, fiber.StatusUnauthorized, "Invalid or missing access token", nil)
	}

	if err := c.BodyParser(&updateUserReq); err != nil {
		return response.JSONErrorResponse(c, fiber.StatusUnauthorized, err.Error(), nil)
	}

	req := domain.UpdateUserByIdRequest{
		AccessToken:      tokenString,
		Email:            updateUserReq.Email,
		Fullname:         updateUserReq.Fullname,
		ProfileImageName: updateUserReq.ProfileImageName,
	}

	result, err := h.svc.UpdateUserById(req)
	if err != nil {
		if validationErrs, ok := err.(domain.ValidationError); ok {
			return response.JSONErrorResponse(c, fiber.StatusBadRequest, "", &validationErrs.Errors)
		}

		return response.JSONErrorResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return response.JSONSuccessResponse(c, result)
}
