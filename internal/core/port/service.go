package port

import (
	"context"

	"github.com/NonthapatKim/many-tooth-api/internal/core/domain"
)

type Service interface {
	// Interest
	CreateUserInterest(req domain.CreateUserInterestRequest) (domain.CreateUserInterestResponse, error)
	GetInterests() ([]domain.GetInterestsResponse, error)

	// Product
	GetProducts() ([]domain.GetProductsResponse, error)
	GetProductCategories() ([]domain.GetProductCategoriesResponse, error)
	GetProductType() ([]domain.GetProuctTypeResponse, error)

	// User
	CheckUser(req domain.CheckUserRequest) (domain.CheckUserResult, error)
	GetUserById(req domain.GetUserByIdRequest) (domain.GetUserByIdResponse, error)
	UserLogin(req domain.UserLoginRequest) (domain.UserLoginResponse, error)
	UserLoginBySocial(req domain.UserLoginBySocialRequest) (domain.UserLoginBySocialResponse, error)
	UserLogout(req domain.UserLogoutRequest) (domain.UserLogoutResponse, error)
	UserRegister(req domain.UserRegisterRequest) (domain.UserRegisterResponse, error)
	UserRegisterBySocial(req domain.UserRegisterBySocialRequest) (domain.UserRegisterBySocialResponse, error)
	UserRequestResetPassword(req domain.UserRequestResetPasswordRequest) (domain.UserRequestResetPasswordResponse, error)
	UserResetPassword(req domain.UserResetPasswordRequest) (domain.UserResetPasswordResponse, error)

	// Mixed
	UserFavProductById(req domain.UserFavProductByIdRequest) (domain.UserFavProductByIdResponse, error)
	GetUserFavProduct(req domain.GetUserFavProductRequest) ([]domain.GetUserFavProductResponse, error)

	// Other
	UserAuthByLine(ctx context.Context, req domain.UserAuthByLineRequest) (domain.UserAuthByLineResponse, error)
	UserAuthByGoogle(ctx context.Context, req domain.UserAuthByGoogleRequest) (domain.UserAuthByGoogleResponse, error)

	// Refresh Token
	CreateRefreshToken(req domain.CreateRefreshTokenRequest) (domain.CreateRefreshTokenResponse, error)
	GetRefreshToken(req domain.GetRefreshTokenRequest) (domain.GetRefreshTokenResponse, error)
	SaveRefreshToken(req domain.SaveRefreshTokenRequest) error
}
