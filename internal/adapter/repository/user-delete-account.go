package repository

import (
	"fmt"
	"time"

	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
)

func (r *repository) UserDeleteAccount(req domain.UserDeleteAccountRequest) (domain.UserDeleteAccountResponse, error) {
	now := time.Now()

	//Revoked Refresh Token
	_, err := r.db.Exec(
		`UPDATE 
			refresh_tokens
			SET 
				counter = ?,
				revoked_at = ? 
			WHERE 
				user_id = ?
				AND local_device_token = ?
		`,
		0,
		now,
		req.UserId,
		req.LocalDeviceToken,
	)
	if err != nil {
		return domain.UserDeleteAccountResponse{}, fmt.Errorf("failed to revoked refresh token: %w", err)
	}

	// Delete User
	_, err = r.db.Exec(
		`DELETE FROM
			users
		WHERE
			user_id = ?
		`,
		req.UserId,
	)
	if err != nil {
		return domain.UserDeleteAccountResponse{}, fmt.Errorf("failed to delete user: %w", err)
	}

	return domain.UserDeleteAccountResponse{}, nil
}
