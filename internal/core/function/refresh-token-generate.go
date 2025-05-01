package function

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateRefreshToken(jti string, userId string, counter int) (string, error) {
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		return "", errors.New("missing secret key in environment variables")
	}

	jti = fmt.Sprintf("%s.%d", jti, counter)

	claims := jwt.MapClaims{
		"jti":     jti,
		"user_id": userId,
		"iat":     time.Now().Unix(),
		"nbf":     time.Now().Add(time.Second * 10).Unix(),
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
