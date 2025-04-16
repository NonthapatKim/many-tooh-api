package repository

import (
	"fmt"

	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
)

func (r *repository) CheckUser(req domain.CheckUserRequest) (domain.CheckUserResult, error) {
	var exists bool

	query := `
		SELECT EXISTS(SELECT 1 FROM user WHERE email = ?)
	`

	err := r.db.QueryRow(query, req.Email).Scan(&exists)
	if err != nil {
		return domain.CheckUserResult{}, fmt.Errorf("error checking: %w", err)
	}

	return domain.CheckUserResult{
		Exists: exists,
	}, nil
}
