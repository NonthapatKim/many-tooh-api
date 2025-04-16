package domain

type UserLoginBySocialRequest struct {
	Token  string `json:"token"`
	Method string `json:"method" example:"google"`
	Email  string
}

type UserLoginBySocialResult struct {
	UserId string
}

type UserLoginBySocialResponse struct {
	Code        int    `json:"code" example:"200"`
	Message     string `json:"message" example:"successfully login"`
	AccessToken string `json:"access_token" example:"eyJhbGciOiJIUzI1NiI..."`
}
