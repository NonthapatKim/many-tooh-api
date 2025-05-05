package service

import (
	"context"
	"errors"
	"os"

	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
	"github.com/NonthapatKim/many-tooth-api/internal/core/function"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func (s *service) AddProductByUser(req domain.AddProductByUserRequest) (domain.AddProductByUserResponse, error) {
	if req.AccessToken == "" {
		return domain.AddProductByUserResponse{}, errors.New("token is required")
	}

	userId, err := function.ValidateAccessToken(&req.AccessToken)
	if err != nil {
		return domain.AddProductByUserResponse{}, err
	}

	reqUserExists := domain.CheckExistsRequest{
		Table:  "users",
		Column: "user_id",
		Id:     &userId,
	}

	exists, err := s.repo.CheckExists(reqUserExists)
	if err != nil {
		return domain.AddProductByUserResponse{}, err
	}
	if !exists.Exists {
		return domain.AddProductByUserResponse{}, errors.New("error: user not found")
	}

	req.UserId = userId

	CLOUDINARY_URL := os.Getenv("CLOUDINARY_URL")
	cld, err := cloudinary.NewFromURL(CLOUDINARY_URL)
	if err != nil {
		return domain.AddProductByUserResponse{}, errors.New("failed to initialize Cloudinary")
	}

	uploadRes, err := cld.Upload.Upload(context.Background(), req.Image, uploader.UploadParams{
		Folder: "product-by-user",
	})
	if err != nil {
		return domain.AddProductByUserResponse{}, errors.New("failed to upload image to Cloudinary")
	}

	req.ImageUrl = uploadRes.SecureURL

	_, err = s.repo.AddProductByUser(req)
	if err != nil {
		return domain.AddProductByUserResponse{}, err
	}

	response := domain.AddProductByUserResponse{
		Code:    200,
		Message: "successfully created product by user",
	}

	return response, nil
}
