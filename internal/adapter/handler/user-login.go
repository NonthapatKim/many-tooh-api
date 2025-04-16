package handler

import (
	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
	"github.com/NonthapatKim/many-tooth-api/internal/core/domain/response"
	"github.com/gofiber/fiber/v2"
)

// UserLogin godoc
// @Summary User login
// @Description Authenticate user with email and password
// @Tags User
// @Accept  json
// @Produce  json
// @Param user body domain.UserLoginRequest true "User Login Request"
// @Success 200 {object} domain.UserLoginResponse "User logged in successfully"
// @Failure 400 {object} domain.ValidationError "Validation error"
// @Failure 401 {object} domain.ErrorResponseExample401 "Unauthorized error"
// @Failure 500 {object} domain.ErrorResponseExample500 "Internal server error"
// @Router /users/login [post]
func (h *handler) UserLogin(c *fiber.Ctx) error {
	var user domain.UserLoginRequest

	if err := c.BodyParser(&user); err != nil {
		return response.JSONErrorResponse(c, fiber.StatusUnauthorized, err.Error(), nil)
	}

	req := domain.UserLoginRequest{
		Email:            user.Email,
		Password:         user.Password,
		LocalDeviceToken: user.LocalDeviceToken,
	}

	result, err := h.svc.UserLogin(req)
	if err != nil {
		if validationErrs, ok := err.(domain.ValidationError); ok {
			return response.JSONErrorResponse(c, fiber.StatusBadRequest, "", &validationErrs.Errors)
		}

		return response.JSONErrorResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return response.JSONSuccessResponse(c, result)
}
