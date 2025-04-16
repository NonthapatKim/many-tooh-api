package domain

type CreateRefreshTokenRequest struct {
	LocalDeviceToken string `json:"local_device_token" validate:"required"`
	RefreshToken     string
}

type CreateRefreshTokenResponse struct {
	Code         int    `json:"code" example:"200"`
	Message      string `json:"message" example:"successfully request token"`
	AccessToken  string `json:"access_token" example:"eyJhbGciOiJIUzI1NiI..."`
	RefreshToken string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiI..."`
}
