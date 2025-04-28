package service

import "github.com/NonthapatKim/many-tooth-api/internal/core/domain"

func (s *service) GetBrands() ([]domain.GetBrandsResponse, error) {
	result, err := s.repo.GetBrands()
	if err != nil {
		return nil, err
	}

	return result, nil
}
