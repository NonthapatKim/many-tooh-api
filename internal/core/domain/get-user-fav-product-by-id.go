package domain

type GetUserFavProductRequest struct {
	AccessToken string
	UserId      string
}

type GetUserFavProductResponse struct {
	ProductId       string  `json:"product_id"`
	ProductImageUrl *string `json:"product_image_url"`
	ProductName     string  `json:"product_name"`
}
