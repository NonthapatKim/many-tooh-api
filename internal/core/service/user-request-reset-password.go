package service

import (
	"errors"

	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
	"github.com/NonthapatKim/many-tooth-api/internal/core/function"
)

func (s *service) UserRequestResetPassword(req domain.UserRequestResetPasswordRequest) (domain.UserRequestResetPasswordResponse, error) {
	if req.Email == "" {
		return domain.UserRequestResetPasswordResponse{}, errors.New("email is required")
	}

	Token, err := function.GenerateToken()
	if err != nil {
		return domain.UserRequestResetPasswordResponse{}, err
	}

	req.Token = Token

	_, err = s.repo.UserRequestResetPassword(req)
	if err != nil {
		return domain.UserRequestResetPasswordResponse{}, err
	}

	response := domain.UserRequestResetPasswordResponse{
		Code:  200,
		Token: Token,
	}

	return response, nil
}
