package ierr

import (
	"fmt"
	"net/http"
)

type InvalidOTP struct {
	PhoneNumber string `json:"phone_number"`
	OTP         string `json:"otp"`
}

func (u InvalidOTP) Error() string {
	return fmt.Sprintf("invalid OTP code with OTP: %v and phone number: %v", u.OTP, u.PhoneNumber)
}

func (u InvalidOTP) HTTPStatusCode() int {
	return http.StatusBadRequest
}

func (u InvalidOTP) HTTPMessage() string {
	return u.Error()
}
