package handler

import (
	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
	"github.com/NonthapatKim/many-tooth-api/internal/core/domain/response"
	"github.com/gofiber/fiber/v2"
)

// CreateRefreshToken godoc
// @Summary Create Refresh Access and Refresh Token
// @Description Refresh the access token and refresh token using a valid refresh token.
// @Tags User
// @Accept json
// @Produce json
// @Param RefreshToken header string true "Refresh Token Token" example:"Bearer <token>"
// @Param body body domain.CreateRefreshTokenRequest true "Refresh Token Request"
// @Success 200 {object} domain.CreateRefreshTokenResponse
// @Failure 400 {object} domain.ErrorResponseExample400
// @Failure 401 {object} domain.ErrorResponseExample401
// @Failure 500 {object} domain.ErrorResponseExample500
// @Router /user/refresh-token [post]
func (h *handler) CreateRefreshToken(c *fiber.Ctx) error {
	var refreshTokenReq domain.CreateRefreshTokenRequest

	refreshToken := c.Locals("userToken").(string)

	if err := c.BodyParser(&refreshTokenReq); err != nil {
		return response.JSONErrorResponse(c, fiber.StatusUnauthorized, err.Error(), nil)
	}

	req := domain.CreateRefreshTokenRequest{
		RefreshToken:     refreshToken,
		LocalDeviceToken: refreshTokenReq.LocalDeviceToken,
	}

	result, err := h.svc.CreateRefreshToken(req)
	if err != nil {
		if validationErrs, ok := err.(domain.ValidationError); ok {
			return response.JSONErrorResponse(c, fiber.StatusBadRequest, "", &validationErrs.Errors)
		}

		return response.JSONErrorResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return response.JSONSuccessResponse(c, result)
}
