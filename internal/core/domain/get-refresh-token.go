package domain

import "time"

type GetRefreshTokenRequest struct {
	UserId           string
	LocalDeviceToken string
}

type GetRefreshTokenResponse struct {
	Counter   int `default:"1"`
	RevokedAt *time.Time
}
