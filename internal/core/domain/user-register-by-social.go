package domain

type UserRegisterBySocialRequest struct {
	Email string
	Name  string
}

type UserRegisterBySocialResponse struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"success"`
}
