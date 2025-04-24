package handler

import (
	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
	"github.com/NonthapatKim/many-tooth-api/internal/core/domain/response"
	"github.com/gofiber/fiber/v2"
)

// UserLogout godoc
// @Summary User logout
// @Description Logs out And Revoked Refresh Token a user by invalidating the session
// @Tags User
// @Accept json
// @Produce json
// @Param userLogout body domain.UserLogoutRequest true "User logout details"
// @Success 200 {object} domain.UserLogoutResponse
// @Failure 401 {object} domain.ErrorResponseExample401
// @Failure 500 {object} domain.ErrorResponseExample500
// @Router /user/logout [post]
func (h *handler) UserLogout(c *fiber.Ctx) error {
	var userLogout domain.UserLogoutRequest

	accessToken := c.Locals("access_token")
	tokenString, ok := accessToken.(string)
	if !ok {
		return response.JSONErrorResponse(c, fiber.StatusUnauthorized, "Invalid or missing access token", nil)
	}

	if err := c.BodyParser(&userLogout); err != nil {
		return response.JSONErrorResponse(c, fiber.StatusUnauthorized, err.Error(), nil)
	}

	req := domain.UserLogoutRequest{
		LocalDeviceToken: userLogout.LocalDeviceToken,
		AccessToken:      tokenString,
	}

	result, err := h.svc.UserLogout(req)
	if err != nil {
		return response.JSONErrorResponse(c, fiber.StatusUnauthorized, err.Error(), nil)
	}

	return response.JSONSuccessResponse(c, result)
}
