package domain

type GetUserByIdRequest struct {
	AccessToken string
	UserId      string
}

type GetUserByIdResponse struct {
	UserId           string  `json:"user_id"`
	ProfileImageName *string `json:"profile_image_name"`
	Email            string  `json:"email"`
	Fullname         string  `json:"fullname"`
	HasInterest      bool    `json:"has_interest"`
	IsPerry          bool    `json:"is_perry"`
}
