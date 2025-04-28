package repository

import (
	"fmt"

	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
)

func (r *repository) GetBrands() ([]domain.GetBrandsResponse, error) {
	var result []domain.GetBrandsResponse

	query := `
		SELECT
			brand_id,
			name AS brand_name
		FROM brands
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var brand domain.GetBrandsResponse
		err := rows.Scan(
			&brand.BrandId,
			&brand.BrandName,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning data: %w", err)
		}
		result = append(result, brand)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating data: %w", err)
	}

	return result, nil
}
