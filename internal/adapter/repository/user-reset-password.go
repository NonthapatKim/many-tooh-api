package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
)

func (r *repository) UserResetPassword(req domain.UserResetPasswordRequest) (domain.UserResetPasswordResult, error) {
	var result domain.UserResetPasswordResult

	err := r.db.QueryRow(
		`SELECT 
			CASE 
				WHEN EXISTS(SELECT 1 FROM user WHERE reset_password_token = ?) 
				THEN (SELECT user_id FROM user WHERE reset_password_token = ? LIMIT 1) 
				ELSE NULL 
			END AS user_id;
		`,
		req.Token,
		req.Token,
	).Scan(&result.UserId)
	if err != nil {
		if err == sql.ErrNoRows {
			return result, errors.New("token failed: token invalid")
		}
		return result, err
	}

	_, err = r.db.Exec(
		`UPDATE user SET password = ? WHERE user_id = ?`,
		req.Password,
		result.UserId,
	)
	if err != nil {
		return result, fmt.Errorf("failed to update user password: %w", err)
	}

	result.Status = true
	return result, nil
}
