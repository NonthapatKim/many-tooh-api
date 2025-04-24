package repository

import (
	"fmt"

	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
)

func (r *repository) GetUserFavProduct(req domain.GetUserFavProductRequest) ([]domain.GetUserFavProductResponse, error) {
	var result []domain.GetUserFavProductResponse

	query := `
		SELECT
			prod.product_id,
			prod.image_url AS product_image_url,
			prod.name AS product_name
		FROM user_favorite_product user_fav
		INNER JOIN products prod
			ON user_fav.product_id = prod.product_id
		WHERE user_fav.user_id = ?
	`
	rows, err := r.db.Query(query, req.UserId)
	if err != nil {
		return nil, fmt.Errorf("error querying: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var product domain.GetUserFavProductResponse
		err := rows.Scan(
			&product.ProductId,
			&product.ProductName,
			&product.ProductImageUrl,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning data: %w", err)
		}
		result = append(result, product)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating data: %w", err)
	}

	if len(result) == 0 {
		return nil, nil
	}

	return result, nil
}
