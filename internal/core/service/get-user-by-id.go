package service

import (
	"errors"

	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
	"github.com/NonthapatKim/many-tooth-api/internal/core/function"
)

func (s *service) GetUserById(req domain.GetUserByIdRequest) (domain.GetUserByIdResponse, error) {
	if req.AccessToken == "" {
		return domain.GetUserByIdResponse{}, errors.New("token is required")
	}

	userId, err := function.ValidateAccessToken(&req.AccessToken)
	if err != nil {
		return domain.GetUserByIdResponse{}, err
	}

	req.UserId = userId

	result, err := s.repo.GetUserById(req)
	if err != nil {
		return result, err
	}

	return result, nil
}
