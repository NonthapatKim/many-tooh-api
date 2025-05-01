package service

import "github.com/NonthapatKim/many-tooth-api/internal/core/domain"

func (s *service) GetProductType() ([]domain.GetProuctTypeResponse, error) {
	result, err := s.repo.GetProductType()
	if err != nil {
		return nil, err
	}

	return result, nil
}
