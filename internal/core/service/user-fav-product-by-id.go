package service

import (
	"errors"
	"fmt"

	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
	"github.com/NonthapatKim/many-tooth-api/internal/core/function"
)

func (s *service) UserFavProductById(req domain.UserFavProductByIdRequest) (domain.UserFavProductByIdResponse, error) {
	var checkList []domain.CheckExistsRequest

	if req.AccessToken == "" {
		return domain.UserFavProductByIdResponse{}, errors.New("token is required")
	}

	if req.ProductId == "" {
		return domain.UserFavProductByIdResponse{}, errors.New("product_id is required")
	}

	userId, err := function.ValidateAccessToken(&req.AccessToken)
	if err != nil {
		return domain.UserFavProductByIdResponse{}, err
	}

	checkList = []domain.CheckExistsRequest{
		{
			Table:  "users",
			Column: "user_id",
			Id:     &userId,
		},
		{
			Table:  "products",
			Column: "product_id",
			Id:     &req.ProductId,
		},
	}

	for _, check := range checkList {
		exists, err := s.repo.CheckExists(check)
		if err != nil {
			return domain.UserFavProductByIdResponse{}, err
		}
		if !exists.Exists {
			return domain.UserFavProductByIdResponse{}, fmt.Errorf("error: %s not found in %s", check.Column, check.Table)
		}
	}

	req.UserId = userId

	result, err := s.repo.UserFavProductById(req)
	if err != nil {
		return result, err
	}

	return result, nil
}
