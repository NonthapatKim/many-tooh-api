package service

import (
	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
)

func (s *service) GetProductCategories() ([]domain.GetProductCategoriesResponse, error) {
	result, err := s.repo.GetProductCategories()
	if err != nil {
		return nil, err
	}

	return result, nil
}
