package repository

import (
	"fmt"

	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
)

func (r *repository) GetUserById(req domain.GetUserByIdRequest) (domain.GetUserByIdResponse, error) {
	var result domain.GetUserByIdResponse

	query := `
		SELECT 
			u.user_id,
			u.profile_image_name,
			u.email,
			u.fullname,
			CASE 
				WHEN EXISTS (
					SELECT 1 
					FROM user_interest ui 
					WHERE ui.user_id = u.user_id
				) THEN 'true'
				ELSE 'false'
			END AS has_interest,
			u.is_perry
		FROM users u
		WHERE u.user_id = ?
			AND u.user_type = ?;	
	`
	err := r.db.QueryRow(query, req.UserId, "user").Scan(
		&result.UserId,
		&result.ProfileImageName,
		&result.Email,
		&result.Fullname,
		&result.HasInterest,
		&result.IsPerry,
	)
	if err != nil {
		return domain.GetUserByIdResponse{}, fmt.Errorf("error scanning user: %w", err)
	}

	return result, nil
}
