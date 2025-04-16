package repository

import (
	"fmt"
	"time"

	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
)

func (r *repository) UserRegisterBySocial(req domain.UserRegisterBySocialRequest) (domain.UserRegisterBySocialResponse, error) {
	query := `
		INSERT INTO user (
			email, 
			created_at
		) VALUES (?, ?)
	`
	_, err := r.db.Exec(
		query,
		req.Email,
		time.Now(),
	)
	if err != nil {
		return domain.UserRegisterBySocialResponse{}, fmt.Errorf("error creating user: %w", err)
	}

	return domain.UserRegisterBySocialResponse{}, nil
}
