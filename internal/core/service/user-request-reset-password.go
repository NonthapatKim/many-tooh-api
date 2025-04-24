package service

import (
	"errors"

	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
	"github.com/NonthapatKim/many-tooth-api/internal/core/function"
)

func (s *service) UserRequestResetPassword(req domain.UserRequestResetPasswordRequest) (domain.UserRequestResetPasswordResponse, error) {
	var validationErrors domain.ValidationErrorResponse
	if req.Email == "" {
		return domain.UserRequestResetPasswordResponse{}, errors.New("email is required")
	}

	reqUserExists := domain.CheckExistsRequest{
		Table:  "users",
		Column: "email",
		Id:     &req.Email,
	}

	exists, err := s.repo.CheckExists(reqUserExists)
	if err != nil {
		return domain.UserRequestResetPasswordResponse{}, err
	}
	if !exists.Exists {
		message := "ขออภัย ไม่พบผู้ใช้งานที่ใช้อีเมลนี้ กรุณาลองใหม่อีกครั้ง"
		validationErrors.UserError = &message
		return domain.UserRequestResetPasswordResponse{}, domain.ValidationError{Errors: validationErrors}
	}

	Token, err := function.GenerateToken()
	if err != nil {
		return domain.UserRequestResetPasswordResponse{}, err
	}

	req.Token = Token

	_, err = s.repo.UserRequestResetPassword(req)
	if err != nil {
		return domain.UserRequestResetPasswordResponse{}, err
	}

	response := domain.UserRequestResetPasswordResponse{
		Code:  200,
		Token: Token,
	}

	return response, nil
}
