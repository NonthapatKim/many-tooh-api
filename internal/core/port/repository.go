package port

import "github.com/NonthapatKim/many-tooth-api/internal/core/domain"

type Repository interface {
	// Interest
	GetInterests() ([]domain.GetInterestsResponse, error)

	// User
	CheckUser(req domain.CheckUserRequest) (domain.CheckUserResult, error)
	UserLogin(req domain.UserLoginRequest) (domain.UserLoginResult, error)
	UserLoginBySocial(req domain.UserLoginBySocialRequest) (domain.UserLoginBySocialResult, error)
	UserLogout(req domain.UserLogoutRequest) (domain.UserLogoutResponse, error)
	UserRegister(req domain.UserRegisterRequest) (domain.UserRegisterResponse, error)
	UserRegisterBySocial(req domain.UserRegisterBySocialRequest) (domain.UserRegisterBySocialResponse, error)
	UserRequestResetPassword(req domain.UserRequestResetPasswordRequest) (domain.UserRequestResetPasswordResult, error)
	UserResetPassword(req domain.UserResetPasswordRequest) (domain.UserResetPasswordResult, error)

	// Refresh Token
	GetRefreshToken(req domain.GetRefreshTokenRequest) (domain.GetRefreshTokenResponse, error)
	SaveRefreshToken(req domain.SaveRefreshTokenRequest) error
}
