package ierr

type HTTPErrorHandler interface {
	HTTPStatusCode() int
	HTTPMessage() string
}
