package domain

import "time"

type GetProductsResponse struct {
	ProductId           string     `json:"product_id"`
	BrandId             string     `json:"brand_id"`
	ProductCategoryId   string     `json:"product_category_id"`
	ProductTypeId       string     `json:"product_type_id"`
	BrandName           string     `json:"brand_name"`
	ProductCategoryName string     `json:"product_category_name"`
	ProductTypeName     string     `json:"product_type_name"`
	ProductImageUrl     *string    `json:"product_image_url"`
	ProductName         string     `json:"product_name"`
	Barcode             string     `json:"barcode"`
	Warning             *string    `json:"warning"`
	UsageDescription    *string    `json:"usage_description"`
	AmountFluoride      *string    `json:"amount_fluoride"`
	Properties          *string    `json:"properties"`
	ActiveIngredient    *string    `json:"active_ingredient"`
	DangerousIngredient *string    `json:"dangerous_ingredient"`
	IsDangerous         bool       `json:"is_dangerous"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           *time.Time `json:"updated_at"`
}
