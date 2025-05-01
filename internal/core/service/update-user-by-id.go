package service

import (
	"errors"

	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
	"github.com/NonthapatKim/many-tooth-api/internal/core/function"
)

func (s *service) UpdateUserById(req domain.UpdateUserByIdRequest) (domain.UpdateUserByIdResponse, error) {
	if req.AccessToken == "" {
		return domain.UpdateUserByIdResponse{}, errors.New("token is required")
	}

	userId, err := function.ValidateAccessToken(&req.AccessToken)
	if err != nil {
		return domain.UpdateUserByIdResponse{}, err
	}

	reqUserExists := domain.CheckExistsRequest{
		Table:  "users",
		Column: "user_id",
		Id:     &userId,
	}

	exists, err := s.repo.CheckExists(reqUserExists)
	if err != nil {
		return domain.UpdateUserByIdResponse{}, err
	}
	if !exists.Exists {
		return domain.UpdateUserByIdResponse{}, errors.New("error: user not found")
	}

	req.UserId = userId

	result, err := s.repo.UpdateUserById(req)
	if err != nil {
		return result, err
	}

	return result, nil
}
