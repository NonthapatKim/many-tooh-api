package repository

import (
	"fmt"

	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
)

func (r *repository) GetProductType() ([]domain.GetProuctTypeResponse, error) {
	var result []domain.GetProuctTypeResponse

	query := `
		SELECT
			product_type_id,
			name AS product_type_name
		FROM product_type
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying data: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var productType domain.GetProuctTypeResponse
		err := rows.Scan(
			&productType.ProductTypeId,
			&productType.ProductTypeName,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning data: %w", err)
		}
		result = append(result, productType)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating data: %w", err)
	}

	return result, nil
}
