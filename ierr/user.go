package ierr

import (
	"fmt"
	"net/http"
)

type UserNotFound struct {
	PhoneNumber string `json:"phone_number"`
}

func (u UserNotFound) Error() string {
	return fmt.Sprintf("user with phone number %s not found", u.PhoneNumber)
}

func (u UserNotFound) HTTPStatusCode() int {
	return http.StatusNotFound
}

func (u UserNotFound) HTTPMessage() string {
	return u.Error()
}

type JWTError struct{}

func (u JWTError) Error() string {
	return "invalid or expired credential"
}

func (u JWTError) HTTPMessage() string {
	return u.Error()
}

func (u JWTError) HTTPStatusCode() int {
	return http.StatusUnauthorized
}
