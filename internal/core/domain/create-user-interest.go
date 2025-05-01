package domain

type CreateUserInterestRequest struct {
	AccessToken string
	UserId      string
	InterestId  []string `json:"interest_id"`
}

type CreateUserInterestResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
