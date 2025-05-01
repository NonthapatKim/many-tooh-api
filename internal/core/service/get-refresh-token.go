package service

import (
	"errors"

	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
)

func (s *service) GetRefreshToken(req domain.GetRefreshTokenRequest) (domain.GetRefreshTokenResponse, error) {
	if req.UserId == "" {
		return domain.GetRefreshTokenResponse{}, errors.New("user_id is required")
	}

	if req.LocalDeviceToken == "" {
		return domain.GetRefreshTokenResponse{}, errors.New("local_device_token is required")
	}

	result, err := s.repo.GetRefreshToken(req)
	if err != nil {
		return result, err
	}

	return result, nil
}
