package domain

type UserRegisterRequest struct {
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Password string `json:"password"`
}

type UserRegisterResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
