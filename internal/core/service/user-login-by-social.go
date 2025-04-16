package service

import (
	"context"
	"errors"

	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
	"github.com/NonthapatKim/many-tooth-api/internal/core/function"
)

func (s *service) UserLoginBySocial(req domain.UserLoginBySocialRequest) (domain.UserLoginBySocialResponse, error) {
	var response domain.UserLoginBySocialResponse
	var email string
	var err error

	ctx := context.Background()

	switch req.Method {
	case "google":
		googleRes, err := s.UserAuthByGoogle(ctx, domain.UserAuthByGoogleRequest{Token: req.Token})
		if err != nil {
			return response, err
		}
		email = googleRes.Email
	case "line":
		lineRes, err := s.UserAuthByLine(ctx, domain.UserAuthByLineRequest{Token: req.Token})
		if err != nil {
			return response, err
		}
		email = lineRes.Email
	default:
		return response, errors.New("unsupported login method")
	}

	if email == "" {
		return response, errors.New("email not found")
	}

	req.Email = email

	reqCheckUser := domain.CheckUserRequest{
		Email: req.Email,
	}

	userExists, err := s.repo.CheckUser(reqCheckUser)
	if err != nil {
		return domain.UserLoginBySocialResponse{}, err
	}

	if !userExists.Exists {
		reqCreateUser := domain.UserRegisterBySocialRequest{
			Email: req.Email,
		}

		_, err := s.repo.UserRegisterBySocial(reqCreateUser)
		if err != nil {
			return response, err
		}
	}

	loginSocialResult, err := s.repo.UserLoginBySocial(req)
	if err != nil {
		return response, err
	}

	accessToken, err := function.GenerateAccessToken(loginSocialResult.UserId)
	if err != nil {
		return response, errors.New("error generating token")
	}

	response = domain.UserLoginBySocialResponse{
		Code:        200,
		Message:     "successfully login",
		AccessToken: accessToken,
	}

	return response, nil

}
