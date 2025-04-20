package service

import (
	"errors"
	"fmt"

	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
	"github.com/NonthapatKim/many-tooth-api/internal/core/function"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) UserResetPassword(req domain.UserResetPasswordRequest) (domain.UserResetPasswordResponse, error) {
	var validationErrors domain.ValidationErrorResponse
	var response domain.UserResetPasswordResponse
	var message string

	if req.Token == "" {
		return response, errors.New("token is required")
	}

	if isValid, err := function.ValidateToken(req.Token); err != nil || !isValid {
		return response, fmt.Errorf("invalid token: %w", err)
	}

	if req.Password == "" {
		message = "กรุณากรอกรหัสผ่านของคุณ"
		validationErrors.Password = &message
	}

	if (validationErrors != domain.ValidationErrorResponse{}) {
		return response, domain.ValidationError{Errors: validationErrors}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return response, err
	}

	req.Password = string(hashedPassword)

	_, err = s.repo.UserResetPassword(req)
	if err != nil {
		return response, err
	}

	response = domain.UserResetPasswordResponse{
		Code:    200,
		Message: "successfully reset password",
	}

	return response, nil
}
