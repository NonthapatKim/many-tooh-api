package repository

import (
	"fmt"

	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
)

func (r *repository) CreateUserInterest(req domain.CreateUserInterestRequest) (domain.CreateUserInterestResponse, error) {
	for _, interest := range req.InterestId {
		query := `
			INSERT INTO user_interest (
				user_id,
				interest_id
			) VALUES (?, ?)
		`
		_, err := r.db.Exec(
			query,
			req.UserId,
			interest,
		)
		if err != nil {
			return domain.CreateUserInterestResponse{}, fmt.Errorf("error creating: %w", err)
		}
	}

	return domain.CreateUserInterestResponse{}, nil
}
