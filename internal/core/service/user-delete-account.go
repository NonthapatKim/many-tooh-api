package service

import (
	"errors"
	"fmt"

	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
	"github.com/NonthapatKim/many-tooth-api/internal/core/function"
)

func (s *service) UserDeleteAccount(req domain.UserDeleteAccountRequest) (domain.UserDeleteAccountResponse, error) {
	if req.AccessToken == "" {
		return domain.UserDeleteAccountResponse{}, errors.New("token is required")
	}

	if req.LocalDeviceToken == "" {
		return domain.UserDeleteAccountResponse{}, fmt.Errorf("local_device_token is required")
	}

	userId, err := function.ValidateAccessToken(&req.AccessToken)
	if err != nil {
		return domain.UserDeleteAccountResponse{}, err
	}

	req.UserId = userId

	result, err := s.repo.UserDeleteAccount(req)
	if err != nil {
		return result, err
	}

	response := domain.UserDeleteAccountResponse{
		Code:    200,
		Message: "successfully logged out and deleted account",
	}

	return response, nil
}
