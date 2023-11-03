package main

import (
	"fmt"
	"log"
	"rekber/config"
	"rekber/http"
	userHandlerHTTP "rekber/http/user"
	userService "rekber/internal/user"
	"rekber/postgres"
	otpRepository "rekber/postgres/otp"
	userRepository "rekber/postgres/user"
	"strconv"

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
		config.Get().PSQL.Host,
		strconv.Itoa(config.Get().PSQL.Port),
		config.Get().PSQL.UserName,
		config.Get().PSQL.Password,
		config.Get().PSQL.DBName,
		config.Get().PSQL.SSLMode,
	)

	app := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: http.ErrorHandler,
	})
	app.Use(logger.New())
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": "OK",
		})
	})

	api := app.Group("/api")
	v1 := api.Group("/v1")

	httpHandlers := initHTTPHandlers(db)
	for _, v := range httpHandlers {
		v.InitRouter(v1)
	}

	if err := app.Listen(fmt.Sprintf(":%s", config.Get().App.Port)); err != nil {
		log.Fatalf("failed to start server: %v", err.Error())
	}
}
