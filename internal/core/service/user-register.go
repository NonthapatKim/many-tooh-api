package service

import (
	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
	"github.com/NonthapatKim/many-tooth-api/internal/core/function"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	validate.RegisterValidation("customEmail", function.CustomEmailValidator)
	validate.RegisterValidation("customPassword", function.CustomPasswordValidator)
}

func (s *service) UserRegister(req domain.UserRegisterRequest) (domain.UserRegisterResponse, error) {
	var validationErrors domain.ValidationErrorResponse
	var response domain.UserRegisterResponse
	var message string

	err := validate.Struct(req)
	if err != nil {
		validationErrors = ProcessValidationError(err)
		return response, domain.ValidationError{Errors: validationErrors}
	}

	reqCheckUser := domain.CheckUserRequest{
		Email: req.Email,
	}

	userExists, err := s.repo.CheckUser(reqCheckUser)
	if err != nil {
		return domain.UserRegisterResponse{}, err
	}

	if userExists.Exists {
		message = "user already exists"
		response = domain.UserRegisterResponse{
			Code:    409,
			Message: message,
		}
		return response, nil
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return response, err
	}
	hashedPasswordStr := string(hashedPassword)
	req.Password = hashedPasswordStr

	_, err = s.repo.UserRegister(req)
	if err != nil {
		return response, err
	}

	response = domain.UserRegisterResponse{
		Code:    200,
		Message: "successfully created user",
	}

	return response, nil
}
