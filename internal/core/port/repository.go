package port

import "github.com/NonthapatKim/many-tooth-api/internal/core/domain"

type Repository interface {
	// Brand
	GetBrands() ([]domain.GetBrandsResponse, error)

	// Interest
	CreateUserInterest(req domain.CreateUserInterestRequest) (domain.CreateUserInterestResponse, error)
	GetInterests() ([]domain.GetInterestsResponse, error)

	// Product
	GetProducts(req domain.GetProductsRequest) ([]domain.GetProductsResponse, error)
	GetProductCategories() ([]domain.GetProductCategoriesResponse, error)
	GetProductType() ([]domain.GetProuctTypeResponse, error)

	// User
	CheckUser(req domain.CheckUserRequest) (domain.CheckUserResult, error)
	GetUserById(req domain.GetUserByIdRequest) (domain.GetUserByIdResponse, error)
	UpdateUserById(req domain.UpdateUserByIdRequest) (domain.UpdateUserByIdResponse, error)
	UserLogin(req domain.UserLoginRequest) (domain.UserLoginResult, error)
	UserLoginBySocial(req domain.UserLoginBySocialRequest) (domain.UserLoginBySocialResult, error)
	UserLogout(req domain.UserLogoutRequest) (domain.UserLogoutResponse, error)
	UserRegister(req domain.UserRegisterRequest) (domain.UserRegisterResponse, error)
	UserRegisterBySocial(req domain.UserRegisterBySocialRequest) (domain.UserRegisterBySocialResponse, error)
	UserRequestResetPassword(req domain.UserRequestResetPasswordRequest) (domain.UserRequestResetPasswordResult, error)
	UserResetPassword(req domain.UserResetPasswordRequest) (domain.UserResetPasswordResult, error)

	// Mixed
	UserFavProductById(req domain.UserFavProductByIdRequest) (domain.UserFavProductByIdResponse, error)
	GetUserFavProduct(req domain.GetUserFavProductRequest) ([]domain.GetUserFavProductResponse, error)
	GetProductByInterest(req domain.GetProductByInterestRequest) ([]domain.GetProductsResponse, error)

	// Refresh Token
	GetRefreshToken(req domain.GetRefreshTokenRequest) (domain.GetRefreshTokenResponse, error)
	SaveRefreshToken(req domain.SaveRefreshTokenRequest) error

	CheckExists(req domain.CheckExistsRequest) (domain.CheckExistsResponse, error)
}
