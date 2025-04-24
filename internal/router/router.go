package router

import (
	"fmt"
	"os"

	"github.com/NonthapatKim/many-tooth-api/internal/adapter/handler"
	"github.com/NonthapatKim/many-tooth-api/internal/adapter/handler/middleware"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Router struct {
	app *fiber.App
}

const serviceBaseURL = "/api"

func NewRouter(h handler.Handler) (*Router, error) {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	basePath := app.Group(serviceBaseURL)
	basePathV1 := basePath.Group("/v1").Use(middleware.LoggerMiddleware())

	basePathV1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("สวัสดี !")
	})

	// Swagger UI
	basePathV1.Get("/swagger/*", swagger.HandlerDefault)

	basePathV1.Get("/interests", h.GetInterests)

	product := basePathV1.Group("/products")
	{
		product.Get("/", h.GetProducts)
		product.Get("/categories", h.GetProductCategories)
		product.Get("/types", h.GetProductType)
	}

	user := basePathV1.Group("/users")
	{
		user.Get("/", middleware.Authorization(), h.GetUserById)
		user.Post("/login", h.UserLogin)
		user.Post("/login-social", h.UserLoginBySocial)
		user.Post("/logout", middleware.Authorization(), h.UserLogout)
		user.Post("/register", h.UserRegister)
		user.Post("/request-reset-password", h.UserRequestResetPassword)
		user.Put("/reset-password", h.UserResetPassword)

		user.Post("/interests", middleware.Authorization(), h.CreateUserInterest)

		// Product
		user.Get("/fav", middleware.Authorization(), h.GetUserFavProduct)
		user.Put("/fav/:productId", middleware.Authorization(), h.UserFavProductById)
	}

	auth := basePathV1.Group("/auth")
	{
		auth.Post("/refresh", h.CreateRefreshToken)
	}

	return &Router{app: app}, nil
}

func (r *Router) Start() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4200"
	}

	fmt.Println("Listening on port", port)
	return r.app.Listen(":" + port)
}
