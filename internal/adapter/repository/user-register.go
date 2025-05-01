package repository

import (
	"fmt"

	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
)

func (r *repository) UserRegister(req domain.UserRegisterRequest) (domain.UserRegisterResponse, error) {
	query := `
		INSERT INTO users (
			email, 
			fullname,
			password
		) VALUES (?, ?, ?)
	`
	_, err := r.db.Exec(
		query,
		req.Email,
		req.Fullname,
		req.Password,
	)
	if err != nil {
		return domain.UserRegisterResponse{}, fmt.Errorf("error creating user: %w", err)
	}

	return domain.UserRegisterResponse{}, nil
}
