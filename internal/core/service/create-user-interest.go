package service

import (
	"errors"

	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
	"github.com/NonthapatKim/many-tooth-api/internal/core/function"
)

func (s *service) CreateUserInterest(req domain.CreateUserInterestRequest) (domain.CreateUserInterestResponse, error) {
	if req.AccessToken == "" {
		return domain.CreateUserInterestResponse{}, errors.New("token is required")
	}

	userId, err := function.ValidateAccessToken(&req.AccessToken)
	if err != nil {
		return domain.CreateUserInterestResponse{}, err
	}

	reqUserExists := domain.CheckExistsRequest{
		Table:  "users",
		Column: "user_id",
		Id:     &userId,
	}

	exists, err := s.repo.CheckExists(reqUserExists)
	if err != nil {
		return domain.CreateUserInterestResponse{}, err
	}
	if !exists.Exists {
		return domain.CreateUserInterestResponse{}, errors.New("error: user not found")
	}

	req.UserId = userId

	_, err = s.repo.CreateUserInterest(req)
	if err != nil {
		return domain.CreateUserInterestResponse{}, err
	}

	response := domain.CreateUserInterestResponse{
		Code:    200,
		Message: "successfully created schedule",
	}

	return response, nil
}
