package repository

import (
	"fmt"

	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
)

func (r *repository) SaveRefreshToken(req domain.SaveRefreshTokenRequest) error {
	query := `
		INSERT INTO refresh_tokens (
			user_id, 
			local_device_token, 
			last_jti, 
			counter
		) 
		VALUES (?, ?, ?, ?) 
		ON DUPLICATE KEY UPDATE 
			counter = counter + 1, 
			last_jti = ?,
			revoked_at = ?
	`
	_, err := r.db.Exec(
		query,
		&req.UserId,
		&req.LocalDeviceToken,
		&req.Jti,
		&req.Counter,
		&req.Jti,
		nil,
	)

	if err != nil {
		return fmt.Errorf("error creating refresh_token: %w", err)
	}

	return nil
}
