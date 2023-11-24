package middleware

import (
	"rekber/ierr"
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

	t := token{
		accessToken: receiveToken[1],
	}

	userData, err := t.parseAccessToken()
	if err != nil {
		return err
	}

	c.Locals("user-data", userData)
	return c.Next()
}
