package domain

import "mime/multipart"

type AddProductByUserRequest struct {
	AccessToken string
	Image       multipart.File
	ImageUrl    string
	ProductName string `json:"product_name"`
	UserId      string
}

type AddProductByUserResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
