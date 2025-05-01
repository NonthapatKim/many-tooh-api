package handler

import (
	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
	"github.com/NonthapatKim/many-tooth-api/internal/core/domain/response"
	"github.com/gofiber/fiber/v2"
)

// UserResetPassword godoc
// @Summary Reset user password
// @Description Reset the password for the user using a reset password token
// @Tags User
// @Accept json
// @Produce json
// @Param userResetPwd body domain.UserResetPasswordRequest true "User reset password request"
// @Success 200 {object} domain.UserResetPasswordResponse
// @Failure 401 {object} domain.ErrorResponseExample401
// @Failure 500 {object} domain.ErrorResponseExample500
// @Router /users/reset-password [put]
func (h *handler) UserResetPassword(c *fiber.Ctx) error {
	var userResetPwd domain.UserResetPasswordRequest

	if err := c.BodyParser(&userResetPwd); err != nil {
		return response.JSONErrorResponse(c, fiber.StatusUnauthorized, err.Error(), nil)
	}

	req := domain.UserResetPasswordRequest{
		Token:    userResetPwd.Token,
		Password: userResetPwd.Password,
	}

	result, err := h.svc.UserResetPassword(req)
	if err != nil {
		return response.JSONErrorResponse(c, fiber.StatusUnauthorized, err.Error(), nil)
	}

	return response.JSONSuccessResponse(c, result)
}
