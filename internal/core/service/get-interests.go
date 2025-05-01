package service

import "github.com/NonthapatKim/many-tooth-api/internal/core/domain"

func (s *service) GetInterests() ([]domain.GetInterestsResponse, error) {
	result, err := s.repo.GetInterests()
	if err != nil {
		return nil, err
	}

	return result, nil
}
