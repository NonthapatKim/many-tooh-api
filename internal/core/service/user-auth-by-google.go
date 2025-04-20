package service

import (
	"context"
	"errors"
	"os"

	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
	"google.golang.org/api/idtoken"
)

func (s *service) UserAuthByGoogle(ctx context.Context, req domain.UserAuthByGoogleRequest) (domain.UserAuthByGoogleResponse, error) {
	googleClientId := os.Getenv("GOOGLE_CLIENT_ID")

	if req.Token == "" {
		return domain.UserAuthByGoogleResponse{}, errors.New("token is required")
	}

	payload, err := idtoken.Validate(ctx, req.Token, googleClientId)
	if err != nil {
		return domain.UserAuthByGoogleResponse{}, err
	}

	response := domain.UserAuthByGoogleResponse{
		Iss:           payload.Issuer,
		Sub:           payload.Subject,
		Aud:           payload.Audience,
		Email:         getString(payload.Claims["email"]),
		EmailVerified: getString(payload.Claims["email_verified"]),
		Name:          getString(payload.Claims["name"]),
		Picture:       getString(payload.Claims["picture"]),
		GivenName:     getString(payload.Claims["given_name"]),
		FamilyName:    getString(payload.Claims["family_name"]),
		Locale:        getString(payload.Claims["locale"]),
	}

	return response, nil
}

func getString(v interface{}) string {
	if s, ok := v.(string); ok {
		return s
	}
	return ""
}
