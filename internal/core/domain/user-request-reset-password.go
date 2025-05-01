package domain

type UserRequestResetPasswordRequest struct {
	Email string `json:"email" validate:"required"`
	Token string
}

type UserRequestResetPasswordResult struct {
	UserId string `json:"user_id"`
}

type UserRequestResetPasswordResponse struct {
	Code  int    `json:"code" example:"200"`
	Token string `json:"token"`
}
