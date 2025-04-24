package service

import "github.com/NonthapatKim/many-tooth-api/internal/core/domain"

func (s *service) GetProducts() ([]domain.GetProductsResponse, error) {
	result, err := s.repo.GetProducts()
	if err != nil {
		return nil, err
	}

	return result, nil
}
