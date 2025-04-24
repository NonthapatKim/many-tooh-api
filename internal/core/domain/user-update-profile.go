package domain

type UserUpdateProfileRequest struct {
	AccessToken      string
	ProfileImageName *string `json:"profile_image_name"`
	Email            string  `json:"email"`
	Fullname         string  `json:"fullname"`
}
