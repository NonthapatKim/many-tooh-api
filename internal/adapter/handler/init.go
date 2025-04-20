package handler

import (
	"github.com/NonthapatKim/many-tooth-api/internal/core/port"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	svc port.Service
}

type Handler interface {
	// Interest
	GetInterests(c *fiber.Ctx) error

	// User
	UserLogin(c *fiber.Ctx) error
	UserLoginBySocial(c *fiber.Ctx) error
	UserLogout(c *fiber.Ctx) error
	UserRegister(c *fiber.Ctx) error
	UserRequestResetPassword(c *fiber.Ctx) error
	UserResetPassword(c *fiber.Ctx) error

	// Refresh Token
	CreateRefreshToken(c *fiber.Ctx) error
}

func New(svc port.Service) Handler {
	return &handler{svc: svc}
}
