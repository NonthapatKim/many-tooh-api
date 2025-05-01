package domain

type UserResetPasswordRequest struct {
	Token    string `json:"token" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserResetPasswordResult struct {
	Status bool
	UserId string `json:"user_id" example:"1234567890"`
}

type UserResetPasswordResponse struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"successfully reset password"`
}
