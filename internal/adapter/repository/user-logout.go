package repository

import (
	"fmt"
	"time"

	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
)

func (r *repository) UserLogout(req domain.UserLogoutRequest) (domain.UserLogoutResponse, error) {
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
		return domain.UserLogoutResponse{}, fmt.Errorf("failed to revoked refresh token: %w", err)
	}

	return domain.UserLogoutResponse{}, nil
}
