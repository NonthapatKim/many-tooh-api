package repository

import (
	"fmt"

	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
)

func (r *repository) GetProductCategories() ([]domain.GetProductCategoriesResponse, error) {
	var result []domain.GetProductCategoriesResponse

	query := `
		SELECT
			category_id AS product_category_id,
			name AS category_name
		FROM product_categories
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying product: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var productCate domain.GetProductCategoriesResponse
		err := rows.Scan(
			&productCate.ProductCategoryId,
			&productCate.ProductCategoryName,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning data: %w", err)
		}
		result = append(result, productCate)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating data: %w", err)
	}

	return result, nil
}
