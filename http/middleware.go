package http

import (
	"rekber/ierr"
	"rekber/internal/token"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("authorization")
	if authHeader == "" {
		return ierr.AuthorizationHeaderNotFound{}
	}

	receiveToken := strings.Split(authHeader, " ")
	if len(receiveToken) != 2 {
		return ierr.TokenIsNotProvided{}
	}

	if strings.ToLower(receiveToken[0]) != "bearer" {
		return ierr.TokenIsNotProvided{}
	}

	t := token.NewToken(token.WithAccessToken(receiveToken[1]))
	userData, err := t.ParseAccessToken()
	if err != nil {
		return err
	}

	c.Locals("user-data", userData)
	return c.Next()
}
