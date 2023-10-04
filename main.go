package main

import (
	"fmt"
	"log"
	"rekber/config"
	userHandlerHTTP "rekber/http/user"
	userService "rekber/internal/user"
	"rekber/postgres"
	otpRepository "rekber/postgres/otp"
	userRepository "rekber/postgres/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jmoiron/sqlx"
)

type HTTPHandler interface {
	InitRouter(r fiber.Router)
}

func initHTTPHandlers(db *sqlx.DB) []HTTPHandler {
	otpRepo := otpRepository.NewRepository()
	userRepo := userRepository.NewRepository(db)
	userSvc := userService.NewService(userRepo, otpRepo)
	userHandler := userHandlerHTTP.NewHandler(userSvc)

	return []HTTPHandler{
		userHandler,
	}
}

func main() {
	config.SetFromFile("development")

	db := postgres.InitDB(
		config.Get().PSQLHost,
		config.Get().PSQLPort,
		config.Get().PSQLUserName,
		config.Get().PSQLPassword,
		config.Get().PSQLDBName,
		config.Get().PSQLSSLMode,
	)

	app := fiber.New()
	app.Use(logger.New())

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": "OK",
		})
	})

	httpHandlers := initHTTPHandlers(db)
	for _, v := range httpHandlers {
		v.InitRouter(v1)
	}

	if err := app.Listen(fmt.Sprintf(":%s", config.Get().Port)); err != nil {
		log.Fatalf("failed to start server: %v", err.Error())
	}
}
