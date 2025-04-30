package handler

import (
	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
	"github.com/NonthapatKim/many-tooth-api/internal/core/domain/response"
	"github.com/gofiber/fiber/v2"
)

// UserLoginBySocial godoc
// @Summary User login By Social Login (Third Party)
// @Description Authenticates a user with provided credentials by Social Login (Third Party)
// @Tags User
// @Accept json
// @Produce json
// @Param user body domain.UserLoginBySocialRequest true "User login details"
// @Success 200 {object} domain.UserLoginBySocialResponse
// @Failure 400 {object} domain.ErrorResponseExample400
// @Failure 401 {object} domain.ErrorResponseExample401
// @Failure 500 {object} domain.ErrorResponseExample500
// @Router /user/login/social [post]
func (h *handler) UserLoginBySocial(c *fiber.Ctx) error {
	var user domain.UserLoginBySocialRequest

	if err := c.BodyParser(&user); err != nil {
		return response.JSONErrorResponse(c, fiber.StatusUnauthorized, err.Error(), nil)
	}

	req := domain.UserLoginBySocialRequest{
		LocalDeviceToken: user.LocalDeviceToken,
		Method:           user.Method,
		Token:            user.Token,
	}

	result, err := h.svc.UserLoginBySocial(req)
	if err != nil {
		if validationErrs, ok := err.(domain.ValidationError); ok {
			return response.JSONErrorResponse(c, fiber.StatusBadRequest, "", &validationErrs.Errors)
		}

		return response.JSONErrorResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return response.JSONSuccessResponse(c, result)
}
