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

type UserForbiddenAccess struct {
	PhoneNumber string `json:"phone_number"`
}

func (u UserForbiddenAccess) Error() string {
	return fmt.Sprintf("user with phone number %s forbidden to access", u.PhoneNumber)
}

func (u UserForbiddenAccess) HTTPStatusCode() int {
	return http.StatusForbidden
}

func (u UserForbiddenAccess) HTTPMessage() string {
	return u.Error()
}
