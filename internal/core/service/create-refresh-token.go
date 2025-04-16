package service

import (
	"errors"
	"time"

	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
	"github.com/NonthapatKim/many-tooth-api/internal/core/function"
	"github.com/nats-io/nuid"
)

func (s *service) CreateRefreshToken(req domain.CreateRefreshTokenRequest) (domain.CreateRefreshTokenResponse, error) {
	if req.RefreshToken == "" {
		return domain.CreateRefreshTokenResponse{}, errors.New("refresh_token is required")
	}

	claims, err := function.ValidateRefreshToken(req.RefreshToken)
	if err != nil {
		return domain.CreateRefreshTokenResponse{}, errors.New("invalid token")
	}

	userId, ok := claims["user_id"].(string)
	if !ok {
		return domain.CreateRefreshTokenResponse{}, errors.New("user_id not found in token claims")
	}

	nbfFloat, ok := claims["nbf"].(float64)
	if !ok {
		return domain.CreateRefreshTokenResponse{}, errors.New("nbf not found in token claims")
	}

	nbfTime := time.Unix(int64(nbfFloat), 0)
	if time.Now().Before(nbfTime) {
		return domain.CreateRefreshTokenResponse{}, errors.New("refresh token is not valid yet")
	}

	// Re-New Access Token
	accessToken, err := function.GenerateAccessToken(userId)
	if err != nil {
		return domain.CreateRefreshTokenResponse{}, errors.New("error generating token")
	}

	// Fetch Refresh Token Counter
	reqCount := domain.GetRefreshTokenRequest{
		UserId:           userId,
		LocalDeviceToken: req.LocalDeviceToken,
	}

	counterResult, err := s.repo.GetRefreshToken(reqCount)
	if err != nil {
		return domain.CreateRefreshTokenResponse{}, errors.New("error getting refresh token counter")
	}

	if counterResult.RevokedAt != nil {
		return domain.CreateRefreshTokenResponse{}, errors.New("token has been revoked")
	}

	tokenId := nuid.Next()

	// Update Counter Only if NBF is valid
	counter := counterResult.Counter + 1

	reqSaveRefreshToken := domain.SaveRefreshTokenRequest{
		UserId:           userId,
		LocalDeviceToken: req.LocalDeviceToken,
		Jti:              tokenId,
		Counter:          counter,
	}

	err = s.repo.SaveRefreshToken(reqSaveRefreshToken)
	if err != nil {
		return domain.CreateRefreshTokenResponse{}, errors.New("error saving refresh token")
	}

	refreshToken, err := function.GenerateRefreshToken(tokenId, userId, counter)
	if err != nil {
		return domain.CreateRefreshTokenResponse{}, errors.New("error generating refresh token")
	}

	return domain.CreateRefreshTokenResponse{
		Code:         200,
		Message:      "successfully request token",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
