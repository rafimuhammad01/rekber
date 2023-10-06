package ierr

import "net/http"

type HTTPErrorHandler interface {
	HTTPStatusCode() int
	HTTPMessage() string
}

type Unauthorized struct {
	Reason string
}

func (u Unauthorized) Error() string {
	return u.Reason
}

func (u Unauthorized) HTTPStatusCode() int {
	return http.StatusUnauthorized
}

func (u Unauthorized) HTTPMessage() string {
	return u.Error()
}
