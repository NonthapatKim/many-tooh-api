package domain

type UserLoginRequest struct {
	Email            string  `json:"email" validate:"required,customEmail"`
	Password         *string `json:"password" validate:"required,min=8"`
	LocalDeviceToken string  `json:"local_device_token" validate:"required"`
}

type UserLoginResult struct {
	UserId   string  `json:"user_id"`
	Password *string `json:"password"`
}

type UserLoginResponse struct {
	Code         int    `json:"code" example:"200"`
	Message      string `json:"message" example:"successfully login"`
	AccessToken  string `json:"access_token" example:"eyJhbGciOiJIUzI1NiI..."`
	RefreshToken string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiI..."`
}
