package domain

type UserLoginBySocialRequest struct {
	Name             string
	Token            string `json:"token"`
	LocalDeviceToken string `json:"local_device_token"`
	Method           string `json:"method" example:"google"`
	Email            string
}

type UserLoginBySocialResult struct {
	UserId string `json:"user_id"`
}

type UserLoginBySocialResponse struct {
	Code         int    `json:"code" example:"200"`
	Message      string `json:"message" example:"successfully login"`
	AccessToken  string `json:"access_token" example:"eyJhbGciOiJIUzI1NiI..."`
	RefreshToken string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiI..."`
}
