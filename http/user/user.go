package user

import (
	"context"
	"fmt"
	"net/http"
	httpHandler "rekber/http"
	"rekber/internal/dto"
	"rekber/internal/user"
	"rekber/middleware"

	"github.com/gofiber/fiber/v2"
)

type Service interface {
	Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error)
	Register(ctx context.Context, req dto.RegisterRequest) error
}

type Handler struct {
	svc Service
}

func (h Handler) InitRouter(r fiber.Router) {
	userGroup := r.Group("/user")
	userGroup.Post("/login", h.Login)
	userGroup.Post("/register", h.Register)
	userGroup.Get("/restricted", middleware.AuthMiddleware, func(c *fiber.Ctx) error {
		userData := c.Locals("userData-data").(user.User)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"test": "hello" + userData.ID.String(),
		})
	})
}

func (h Handler) Login(c *fiber.Ctx) error {
	var req dto.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return fmt.Errorf("failed to parse body: %w", err)
	}

	token, err := h.svc.Login(c.Context(), req)
	if err != nil {
		return fmt.Errorf("failed when calling user service: %w", err)
	}

	return c.Status(fiber.StatusOK).JSON(httpHandler.JSONResponse{
		Message: "successfully login",
		Data:    token,
	})
}

func (h Handler) Register(c *fiber.Ctx) error {
	var req dto.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return fmt.Errorf("failed to parse body: %w", err)
	}

	if err := h.svc.Register(c.Context(), req); err != nil {
		return fmt.Errorf("failed when calling user service: %w", err)
	}

	return c.Status(http.StatusOK).JSON(httpHandler.JSONResponse{
		Message: "successfully register user",
	})
}

func NewHandler(svc Service) *Handler {
	return &Handler{
		svc: svc,
	}
}
