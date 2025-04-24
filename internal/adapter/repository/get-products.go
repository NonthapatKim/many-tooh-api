package repository

import (
	"fmt"

	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
)

func (r *repository) GetProducts() ([]domain.GetProductsResponse, error) {
	var result []domain.GetProductsResponse

	query := `
		SELECT
			prod.product_id,
			prod.brand_id,
			prod.product_category_id,
			prod.product_type_id,
			brands.name AS brand_name,
			prod_cate.name AS product_category_name,
			prod_type.name AS product_type_name,
			prod.name AS product_name,
			prod.image_url AS product_image_url,
			prod.barcode,
			prod.warning,
			prod.usage_description,
			prod.amount_fluoride,
			prod.properties,
			prod.active_ingredient,
			prod.dangerous_ingredient,
			prod.created_at,
			prod.updated_at
		FROM products prod
		INNER JOIN brands
			ON prod.brand_id = brands.brand_id
		INNER JOIN product_categories prod_cate
			ON prod.product_category_id = prod_cate.category_id
		INNER JOIN product_type prod_type
			ON prod.product_type_id = prod_type.product_type_id
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying product: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var product domain.GetProductsResponse
		err := rows.Scan(
			&product.ProductId,
			&product.BrandId,
			&product.ProductCategoryId,
			&product.ProductTypeId,
			&product.BrandName,
			&product.ProductCategoryName,
			&product.ProductTypeName,
			&product.ProductName,
			&product.ProductImageUrl,
			&product.Barcode,
			&product.Warning,
			&product.UsageDescription,
			&product.AmountFluoride,
			&product.Properties,
			&product.ActiveIngredient,
			&product.DangerousIngredient,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning data: %w", err)
		}
		result = append(result, product)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating data: %w", err)
	}

	return result, nil
}
