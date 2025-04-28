package handler

import (
	"github.com/NonthapatKim/many-tooth-api/internal/core/port"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	svc port.Service
}

type Handler interface {
	// Brand
	GetBrands(c *fiber.Ctx) error

	// Interest
	CreateUserInterest(c *fiber.Ctx) error
	GetInterests(c *fiber.Ctx) error

	// Product
	GetProducts(c *fiber.Ctx) error
	GetProductCategories(c *fiber.Ctx) error
	GetProductType(c *fiber.Ctx) error

	// User
	GetUserById(c *fiber.Ctx) error
	UpdateUserById(c *fiber.Ctx) error
	UserLogin(c *fiber.Ctx) error
	UserLoginBySocial(c *fiber.Ctx) error
	UserLogout(c *fiber.Ctx) error
	UserRegister(c *fiber.Ctx) error
	UserRequestResetPassword(c *fiber.Ctx) error
	UserResetPassword(c *fiber.Ctx) error

	// Mixed
	UserFavProductById(c *fiber.Ctx) error
	GetUserFavProduct(c *fiber.Ctx) error
	GetProductByInterest(c *fiber.Ctx) error

	// Refresh Token
	CreateRefreshToken(c *fiber.Ctx) error
}

func New(svc port.Service) Handler {
	return &handler{svc: svc}
}
