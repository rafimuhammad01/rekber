package http

import (
	"errors"
	"rekber/ierr"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError
	message := "internal server error"

	// Retrieve the custom status code if it's a *fiber.Error
	var e ierr.HTTPErrorHandler
	if errors.As(err, &e) {
		code = e.HTTPStatusCode()
		message = e.HTTPMessage()

		return c.Status(code).JSON(JSONResponse{
			Error: message,
		})
	}

	// Return status code with error message
	return c.Status(code).JSON(JSONResponse{
		Error: message,
	})
}
