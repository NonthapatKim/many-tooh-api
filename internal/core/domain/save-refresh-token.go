package domain

type SaveRefreshTokenRequest struct {
	UserId           string
	LocalDeviceToken string
	Jti              string
	Counter          int
}

type SaveRefreshTokenResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
