package domain

type GetInterestsResponse struct {
	InterestId       string `json:"interest_id"`
	InterestImageUrl string `json:"interest_image_url"`
	InterestName     string `json:"interest_name"`
}
