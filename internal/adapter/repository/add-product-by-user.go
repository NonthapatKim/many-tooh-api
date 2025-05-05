package repository

import (
	"fmt"

	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
)

func (r *repository) AddProductByUser(req domain.AddProductByUserRequest) (domain.AddProductByUserResponse, error) {
	query := `
		INSERT INTO product_by_user (
			user_id,
			image_url,
			product_name
		) VALUES (?, ?, ?)
	`
	_, err := r.db.Exec(
		query,
		req.UserId,
		req.ImageUrl,
		req.ProductName,
	)
	if err != nil {
		return domain.AddProductByUserResponse{}, fmt.Errorf("error creating user: %w", err)
	}

	return domain.AddProductByUserResponse{}, nil
}
