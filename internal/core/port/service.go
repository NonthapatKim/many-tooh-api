package port

import (
	"context"

	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
)

type Service interface {
	// Interest
	GetInterests() ([]domain.GetInterestsResponse, error)

	// User
	CheckUser(req domain.CheckUserRequest) (domain.CheckUserResult, error)
	UserLogin(req domain.UserLoginRequest) (domain.UserLoginResponse, error)
	UserLoginBySocial(req domain.UserLoginBySocialRequest) (domain.UserLoginBySocialResponse, error)
	UserLogout(req domain.UserLogoutRequest) (domain.UserLogoutResponse, error)
	UserRegister(req domain.UserRegisterRequest) (domain.UserRegisterResponse, error)
	UserRegisterBySocial(req domain.UserRegisterBySocialRequest) (domain.UserRegisterBySocialResponse, error)
	UserRequestResetPassword(req domain.UserRequestResetPasswordRequest) (domain.UserRequestResetPasswordResponse, error)
	UserResetPassword(req domain.UserResetPasswordRequest) (domain.UserResetPasswordResponse, error)

	// Other
	UserAuthByLine(ctx context.Context, req domain.UserAuthByLineRequest) (domain.UserAuthByLineResponse, error)
	UserAuthByGoogle(ctx context.Context, req domain.UserAuthByGoogleRequest) (domain.UserAuthByGoogleResponse, error)

	// Refresh Token
	CreateRefreshToken(req domain.CreateRefreshTokenRequest) (domain.CreateRefreshTokenResponse, error)
	GetRefreshToken(req domain.GetRefreshTokenRequest) (domain.GetRefreshTokenResponse, error)
	SaveRefreshToken(req domain.SaveRefreshTokenRequest) error
}
