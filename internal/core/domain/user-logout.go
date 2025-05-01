package domain

type UserLogoutRequest struct {
	AccessToken      string
	LocalDeviceToken string `json:"local_device_token" validate:"required"`
	UserId           string
}

type UserLogoutResponse struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"successfully logged out"`
}
