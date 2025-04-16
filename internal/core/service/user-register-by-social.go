package service

import "github.com/NonthapatKim/many-tooth-api/internal/core/domain"

func (s *service) UserRegisterBySocial(req domain.UserRegisterBySocialRequest) (domain.UserRegisterBySocialResponse, error) {
	var message string

	reqCheckUser := domain.CheckUserRequest(req)

	userExists, err := s.repo.CheckUser(reqCheckUser)
	if err != nil {
		return domain.UserRegisterBySocialResponse{}, err
	}

	if userExists.Exists {
		message = "user already exists"
		response := domain.UserRegisterBySocialResponse{
			Code:    409,
			Message: message,
		}
		return response, nil
	}

	_, err = s.repo.UserRegisterBySocial(req)
	if err != nil {
		return domain.UserRegisterBySocialResponse{}, err
	}

	response := domain.UserRegisterBySocialResponse{
		Code:    200,
		Message: "successfully created user",
	}

	return response, nil
}
