package domain

type UserRegisterBySocialRequest struct {
	Email string `json:"email"`
}

type UserRegisterBySocialResponse struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"success"`
}
