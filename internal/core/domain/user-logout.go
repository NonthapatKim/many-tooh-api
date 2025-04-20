package domain

type UserLogoutRequest struct {
	LocalDeviceToken string `json:"local_device_token" validate:"required"`
	UserId           string
	UserToken        string
}

type UserLogoutResponse struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"successfully logged out"`
}
