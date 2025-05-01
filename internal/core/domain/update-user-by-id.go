package domain

type UpdateUserByIdRequest struct {
	AccessToken      string
	Email            string  `json:"email"`
	Fullname         string  `json:"fullname"`
	ProfileImageName *string `json:"profile_image_name"`
	UserId           string
}

type UpdateUserByIdResponse struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"successfully updated"`
}
