package port

import "github.com/NonthapatKim/many-tooth-api/internal/core/domain"

type Repository interface {
	// Interest
	GetInterests() ([]domain.GetInterestsResponse, error)

	// User
	CheckUser(req domain.CheckUserRequest) (domain.CheckUserResult, error)
	UserLogin(req domain.UserLoginRequest) (domain.UserLoginResult, error)
	UserLoginBySocial(req domain.UserLoginBySocialRequest) (domain.UserLoginBySocialResult, error)
	UserRegister(req domain.UserRegisterRequest) (domain.UserRegisterResponse, error)
	UserRegisterBySocial(req domain.UserRegisterBySocialRequest) (domain.UserRegisterBySocialResponse, error)

	// Refresh Token
	GetRefreshToken(req domain.GetRefreshTokenRequest) (domain.GetRefreshTokenResponse, error)
	SaveRefreshToken(req domain.SaveRefreshTokenRequest) error
}
