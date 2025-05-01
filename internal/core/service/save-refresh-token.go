package service

import (
	"errors"

	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
)

func (s *service) SaveRefreshToken(req domain.SaveRefreshTokenRequest) error {
	if req.UserId == "" {
		return errors.New("user_id is required")
	}

	if req.LocalDeviceToken == "" {
		return errors.New("local_device_token is required")
	}

	if req.Jti == "" {
		return errors.New("jti is required")
	}

	if req.Counter == 0 {
		return errors.New("counter is required")
	}

	err := s.repo.SaveRefreshToken(req)
	if err != nil {
		return err
	}

	return nil
}
