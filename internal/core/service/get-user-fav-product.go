package service

import (
	"errors"

	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
	"github.com/NonthapatKim/many-tooth-api/internal/core/function"
)

func (s *service) GetUserFavProduct(req domain.GetUserFavProductRequest) ([]domain.GetUserFavProductResponse, error) {
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

	result, err := s.repo.GetUserFavProduct(req)
	if err != nil {
		return nil, err
	}

	return result, nil
}
