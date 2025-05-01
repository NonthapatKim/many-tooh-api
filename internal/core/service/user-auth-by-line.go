package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
)

func (s *service) UserAuthByLine(ctx context.Context, req domain.UserAuthByLineRequest) (domain.UserAuthByLineResponse, error) {
	LINE_CLIENT_ID := os.Getenv("LINE_CLIENT_ID")

	if req.Token == "" {
		return domain.UserAuthByLineResponse{}, errors.New("id_token is required")
	}

	formData := url.Values{}
	formData.Set("id_token", req.Token)
	formData.Set("client_id", LINE_CLIENT_ID)

	client := &http.Client{Timeout: 10 * time.Second}
	reqURL := "https://api.line.me/oauth2/v2.1/verify"
	httpReq, err := http.NewRequest("POST", reqURL, strings.NewReader(formData.Encode()))
	if err != nil {
		return domain.UserAuthByLineResponse{}, err
	}

	httpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	httpReq.Header.Set("Accept", "application/json")

	resp, err := client.Do(httpReq)
	if err != nil {
		return domain.UserAuthByLineResponse{}, err
	}
	defer resp.Body.Close()

	fmt.Println(resp)

	if resp.StatusCode != http.StatusOK {
		return domain.UserAuthByLineResponse{}, errors.New("invalid LINE ID token")
	}

	var lineUser domain.UserAuthByLineResponse
	if err := json.NewDecoder(resp.Body).Decode(&lineUser); err != nil {
		return domain.UserAuthByLineResponse{}, err
	}

	return lineUser, nil
}
