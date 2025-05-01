package service

import (
	"errors"

	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
	"github.com/NonthapatKim/many-tooth-api/internal/core/function"
)

func (s *service) GetProductByInterest(req domain.GetProductByInterestRequest) ([]domain.GetProductsResponse, error) {
	if req.AccessToken == "" {
		return nil, errors.New("token is required")
	}

	userId, err := function.ValidateAccessToken(&req.AccessToken)
	if err != nil {
		return nil, err
	}

	reqUserExists := domain.CheckExistsRequest{
		Table:  "users",
		Column: "user_id",
		Id:     &userId,
	}

	exists, err := s.repo.CheckExists(reqUserExists)
	if err != nil {
		return nil, err
	}
	if !exists.Exists {
		return nil, errors.New("error: user not found")
	}

	req.UserId = userId

	result, err := s.repo.GetProductByInterest(req)
	if err != nil {
		return nil, err
	}

	return result, nil
}
