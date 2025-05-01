package service

import (
	"errors"

	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
	"github.com/NonthapatKim/many-tooth-api/internal/core/function"
	"github.com/go-playground/validator/v10"
	"github.com/nats-io/nuid"
	"gorm.io/gorm"
)

var validate = validator.New()

func init() {
	validate.RegisterValidation("customEmail", function.CustomEmailValidator)
}

func (s *service) UserLogin(req domain.UserLoginRequest) (domain.UserLoginResponse, error) {
	var validationErrors domain.ValidationErrorResponse
	var response domain.UserLoginResponse
	var message string

	err := validate.Struct(req)
	if err != nil {
		validationErrors = ProcessValidationError(err)
		return response, domain.ValidationError{Errors: validationErrors}
	}

	loginResult, err := s.repo.UserLogin(req)
	if err != nil {
		if err.Error() == "incorrect password" {
			message = "อีเมลหรือรหัสผ่านไม่ถูกต้อง กรุณาลองใหม่อีกครั้ง"
			validationErrors.Incorrect = &message
			return response, domain.ValidationError{Errors: validationErrors}
		}

		if err.Error() == "user not found" {
			message = "ไม่พบบัญชีผู้ใช้นี้ กรุณาลองใหม่อีกครั้ง"
			validationErrors.UserError = &message
			return response, domain.ValidationError{Errors: validationErrors}
		}

		if err == gorm.ErrRecordNotFound {
			message = "อีเมลหรือรหัสผ่านไม่ถูกต้อง กรุณาลองใหม่อีกครั้ง"
			validationErrors.Incorrect = &message
			return response, domain.ValidationError{Errors: validationErrors}
		}
		return response, err
	}

	accessToken, err := function.GenerateAccessToken(loginResult.UserId)
	if err != nil {
		return response, errors.New("error generating token")
	}

	//Refresh Token
	reqCount := domain.GetRefreshTokenRequest{
		UserId:           loginResult.UserId,
		LocalDeviceToken: req.LocalDeviceToken,
	}

	counterResult, err := s.repo.GetRefreshToken(reqCount)
	if err != nil {
		return response, err
	}

	tokenId := nuid.Next()
	counter := counterResult.Counter

	reqSaveRefreshToken := domain.SaveRefreshTokenRequest{
		UserId:           loginResult.UserId,
		LocalDeviceToken: req.LocalDeviceToken,
		Jti:              tokenId,
		Counter:          counter,
	}

	err = s.repo.SaveRefreshToken(reqSaveRefreshToken)
	if err != nil {
		return response, errors.New("error save refresh token")
	}

	refreshToken, err := function.GenerateRefreshToken(tokenId, loginResult.UserId, counter)
	if err != nil {
		return response, errors.New("error generate refresh token")
	}

	response = domain.UserLoginResponse{
		Code:         200,
		Message:      "successfully login",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return response, nil
}
