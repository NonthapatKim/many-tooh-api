package service

import (
	"errors"
	"fmt"

	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
	"github.com/NonthapatKim/many-tooth-api/internal/core/function"
)

func (s *service) UserLogout(req domain.UserLogoutRequest) (domain.UserLogoutResponse, error) {
	if req.AccessToken == "" {
		return domain.UserLogoutResponse{}, errors.New("token is required")
	}

	if req.LocalDeviceToken == "" {
		return domain.UserLogoutResponse{}, fmt.Errorf("local_device_token is required")
	}

	userId, err := function.ValidateAccessToken(&req.AccessToken)
	if err != nil {
		return domain.UserLogoutResponse{}, err
	}

	req.UserId = userId

	result, err := s.repo.UserLogout(req)
	if err != nil {
		return result, err
	}

	response := domain.UserLogoutResponse{
		Code:    200,
		Message: "successfully logged out",
	}

	return response, nil
}
