package service

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
)

func (s *service) UserAuthByGoogle(ctx context.Context, req domain.UserAuthByGoogleRequest) (domain.UserAuthByGoogleResponse, error) {
	if req.Token == "" {
		return domain.UserAuthByGoogleResponse{}, errors.New("token is required")
	}

	client := &http.Client{Timeout: 10 * time.Second}
	tokenInfoURL := "https://oauth2.googleapis.com/tokeninfo?id_token=" + req.Token

	httpReq, err := http.NewRequest("GET", tokenInfoURL, nil)
	if err != nil {
		return domain.UserAuthByGoogleResponse{}, errors.New("failed to create request")
	}

	resp, err := client.Do(httpReq)
	if err != nil {
		return domain.UserAuthByGoogleResponse{}, errors.New("failed to verify token with Google")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return domain.UserAuthByGoogleResponse{}, errors.New("invalid token")
	}

	var tokenInfo struct {
		Iss           string `json:"iss"`
		Sub           string `json:"sub"`
		Azp           string `json:"azp"`
		Aud           string `json:"aud"`
		Iat           string `json:"iat"`
		Exp           string `json:"exp"`
		Email         string `json:"email"`
		EmailVerified string `json:"email_verified"`
		Name          string `json:"name"`
		Picture       string `json:"picture"`
		GivenName     string `json:"given_name"`
		FamilyName    string `json:"family_name"`
		Locale        string `json:"locale"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&tokenInfo); err != nil {
		return domain.UserAuthByGoogleResponse{}, errors.New("failed to parse token info")
	}

	response := domain.UserAuthByGoogleResponse{
		Iss:           tokenInfo.Iss,
		Sub:           tokenInfo.Sub,
		Azp:           tokenInfo.Azp,
		Aud:           tokenInfo.Aud,
		Iat:           tokenInfo.Iat,
		Exp:           tokenInfo.Exp,
		Email:         tokenInfo.Email,
		EmailVerified: tokenInfo.EmailVerified,
		Name:          tokenInfo.Name,
		Picture:       tokenInfo.Picture,
		GivenName:     tokenInfo.GivenName,
		FamilyName:    tokenInfo.FamilyName,
		Locale:        tokenInfo.Locale,
	}

	return response, nil
}
