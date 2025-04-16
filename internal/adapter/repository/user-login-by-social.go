package repository

import (
	"database/sql"
	"errors"

	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
)

func (r *repository) UserLoginBySocial(req domain.UserLoginBySocialRequest) (domain.UserLoginBySocialResult, error) {
	var result domain.UserLoginBySocialResult

	err := r.db.QueryRow(
		`SELECT 
			user_id
		FROM user 
		WHERE email = ?`,
		req.Email,
	).Scan(
		&result.UserId,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.UserLoginBySocialResult{}, errors.New("user not found")
		}
		return domain.UserLoginBySocialResult{}, err
	}

	return domain.UserLoginBySocialResult{}, nil
}
