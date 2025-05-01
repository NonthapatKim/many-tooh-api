package handler

import (
	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
	"github.com/NonthapatKim/many-tooth-api/internal/core/domain/response"
	"github.com/gofiber/fiber/v2"
)

// UserRequestResetPassword godoc
// @Summary Request password reset
// @Description Sends a password reset request to the specified email
// @Tags User
// @Accept json
// @Produce json
// @Param userRequestResetPwd body domain.UserRequestResetPasswordRequest true "Email for password reset"
// @Success 200 {object} domain.UserRequestResetPasswordResponse
// @Failure 401 {object} domain.ErrorResponseExample401
// @Failure 500 {object} domain.ErrorResponseExample500
// @Router /users/request-reset-password [post]
func (h *handler) UserRequestResetPassword(c *fiber.Ctx) error {
	var userRequestResetPwd domain.UserRequestResetPasswordRequest

	if err := c.BodyParser(&userRequestResetPwd); err != nil {
		return response.JSONErrorResponse(c, fiber.StatusUnauthorized, err.Error(), nil)
	}

	req := domain.UserRequestResetPasswordRequest{
		Email: userRequestResetPwd.Email,
	}

	result, err := h.svc.UserRequestResetPassword(req)
	if err != nil {
		return response.JSONErrorResponse(c, fiber.StatusUnauthorized, err.Error(), nil)
	}

	return response.JSONSuccessResponse(c, result)
}
