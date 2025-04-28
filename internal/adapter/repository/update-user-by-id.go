package repository

import (
	"fmt"

	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
)

func (r *repository) UpdateUserById(req domain.UpdateUserByIdRequest) (domain.UpdateUserByIdResponse, error) {
	query := `
		UPDATE
			users
		SET
			email = ?,
			fullname = ?,
			profile_image_name = ?
		WHERE
			user_id = ?
	`
	_, err := r.db.Exec(
		query,
		req.Email,
		req.Fullname,
		req.ProfileImageName,
		req.UserId,
	)
	if err != nil {
		return domain.UpdateUserByIdResponse{}, fmt.Errorf("error updating: %w", err)
	}

	return domain.UpdateUserByIdResponse{}, nil
}
