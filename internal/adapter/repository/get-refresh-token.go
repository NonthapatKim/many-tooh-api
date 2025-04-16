package repository

import (
	"database/sql"
	"fmt"

	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
)

func (r *repository) GetRefreshToken(req domain.GetRefreshTokenRequest) (domain.GetRefreshTokenResponse, error) {
	var result domain.GetRefreshTokenResponse
	query := `
		SELECT
			counter,
			revoked_at
		FROM refresh_tokens
		WHERE user_id = ?
			AND local_device_token = ?
	`
	err := r.db.QueryRow(query, req.UserId, req.LocalDeviceToken).Scan(
		&result.Counter,
		&result.RevokedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			result.Counter = 1
			return result, nil
		}
		return domain.GetRefreshTokenResponse{}, fmt.Errorf("error fetching counter: %w", err)
	}

	return result, nil
}
