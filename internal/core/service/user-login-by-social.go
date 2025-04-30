package service

import (
	"context"
	"errors"

	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
	"github.com/NonthapatKim/many-tooth-api/internal/core/function"
	"github.com/nats-io/nuid"
)

func (s *service) UserLoginBySocial(req domain.UserLoginBySocialRequest) (domain.UserLoginBySocialResponse, error) {
	var response domain.UserLoginBySocialResponse
	var email string
	var name string
	var err error

	ctx := context.Background()

	switch req.Method {
	case "google":
		googleRes, err := s.UserAuthByGoogle(ctx, domain.UserAuthByGoogleRequest{Token: req.Token})
		if err != nil {
			return response, err
		}
		email = googleRes.Email
		name = googleRes.Name
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
			Name:  name,
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

	userId := loginSocialResult.UserId

	accessToken, err := function.GenerateAccessToken(loginSocialResult.UserId)
	if err != nil {
		return response, errors.New("error generating token")
	}

	reqCount := domain.GetRefreshTokenRequest{
		UserId:           userId,
		LocalDeviceToken: req.LocalDeviceToken,
	}

	counterResult, err := s.repo.GetRefreshToken(reqCount)
	if err != nil {
		return response, errors.New("error get count refresh token")
	}

	tokenId := nuid.Next()
	counter := counterResult.Counter

	reqSaveRefreshToken := domain.SaveRefreshTokenRequest{
		UserId:           userId,
		LocalDeviceToken: req.LocalDeviceToken,
		Jti:              tokenId,
		Counter:          counter,
	}

	err = s.repo.SaveRefreshToken(reqSaveRefreshToken)
	if err != nil {
		return response, errors.New("error save refresh token")
	}

	refreshToken, err := function.GenerateRefreshToken(tokenId, loginSocialResult.UserId, counter)
	if err != nil {
		return response, errors.New("error generate refresh token")
	}

	response = domain.UserLoginBySocialResponse{
		Code:         200,
		Message:      "successfully login",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return response, nil

}
