package domain

type UserFavProductByIdRequest struct {
	AccessToken string
	UserId      string
	ProductId   string `json:"product_id"`
}

type UserFavProductByIdResponse struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"updated followed success"`
}
