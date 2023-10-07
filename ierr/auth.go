package ierr

import "net/http"

type AuthorizationHeaderNotFound struct{}

func (u AuthorizationHeaderNotFound) Error() string {
	return "authorization header not found"
}

func (u AuthorizationHeaderNotFound) HTTPStatusCode() int {
	return http.StatusUnauthorized
}

func (u AuthorizationHeaderNotFound) HTTPMessage() string {
	return u.Error()
}

type TokenIsNotProvided struct{}

func (u TokenIsNotProvided) Error() string {
	return "token is not provided"
}

func (u TokenIsNotProvided) HTTPStatusCode() int {
	return http.StatusUnauthorized
}

func (u TokenIsNotProvided) HTTPMessage() string {
	return u.Error()
}
