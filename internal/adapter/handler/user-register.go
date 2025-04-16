package handler

import (
	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
	"github.com/NonthapatKim/many-tooth-api/internal/core/domain/response"
	"github.com/gofiber/fiber/v2"
)

// UserRegister godoc
// @Summary Register a new user
// @Description Register a new user with email and password
// @Tags User
// @Accept  json
// @Produce  json
// @Param user body domain.UserRegisterRequest true "User Register Request"
// @Success 200 {object} domain.UserRegisterResponse "User registered successfully"
// @Failure 400 {object} domain.ValidationError "Validation error"
// @Failure 401 {object} domain.ErrorResponseExample401 "Unauthorized error"
// @Failure 500 {object} domain.ErrorResponseExample500 "Internal server error"
// @Router /users/register [post]
func (h *handler) UserRegister(c *fiber.Ctx) error {
	var user domain.UserRegisterRequest

	if err := c.BodyParser(&user); err != nil {
		return response.JSONErrorResponse(c, fiber.StatusUnauthorized, err.Error(), nil)
	}

	req := domain.UserRegisterRequest{
		Email:    user.Email,
		Fullname: user.Fullname,
		Password: user.Password,
	}

	result, err := h.svc.UserRegister(req)
	if err != nil {
		if validationErrs, ok := err.(domain.ValidationError); ok {
			return response.JSONErrorResponse(c, fiber.StatusBadRequest, "", &validationErrs.Errors)
		}

		return response.JSONErrorResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return response.JSONSuccessResponse(c, result)
}
