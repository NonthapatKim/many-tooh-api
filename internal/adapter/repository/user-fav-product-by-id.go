package repository

import (
	"fmt"

	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
)

func (r *repository) UserFavProductById(req domain.UserFavProductByIdRequest) (domain.UserFavProductByIdResponse, error) {
	var exists bool

	err := r.db.QueryRow(`
		SELECT EXISTS (
			SELECT 1
			FROM user_favorite_product
			WHERE user_id = ? AND product_id = ?
		)
	`, req.UserId, req.ProductId).Scan(&exists)

	if err != nil {
		return domain.UserFavProductByIdResponse{}, fmt.Errorf("error checking existence: %w", err)
	}

	if !exists {
		query := `
			INSERT INTO user_favorite_product (
				user_id, 
				product_id
			) VALUES (?, ?)
		`
		_, err := r.db.Exec(query, req.UserId, req.ProductId)
		if err != nil {
			return domain.UserFavProductByIdResponse{}, fmt.Errorf("error inserting data: %w", err)
		}

		return domain.UserFavProductByIdResponse{
			Code:    200,
			Message: "Successfully fav product",
		}, nil
	}

	// Unfav
	query := `
		DELETE 
		FROM 
			user_favorite_product
		WHERE 
			user_id = ? 
			AND product_id = ?
	`
	_, err = r.db.Exec(
		query,
		req.UserId,
		req.ProductId,
	)

	if err != nil {
		return domain.UserFavProductByIdResponse{}, fmt.Errorf("error deleting data: %w", err)
	}

	return domain.UserFavProductByIdResponse{
		Code:    200,
		Message: "Successfully unfav product",
	}, nil
}
