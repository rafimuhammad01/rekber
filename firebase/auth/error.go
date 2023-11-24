package auth

import (
	"net/http"
	"rekber/ierr"
	"strings"
)

const (
	InvalidCode = "INVALID_CODE"
)

func getInvalidCodeError(err interface{}, phoneNumber, otp string) (ierr.InvalidOTP, bool) {
	firebaseErr, ok := err.(Error)
	if !ok {
		return ierr.InvalidOTP{}, false
	}

	if firebaseErr.Code != http.StatusBadRequest {
		return ierr.InvalidOTP{}, false
	}

	for _, v := range firebaseErr.Errors {
		if strings.Contains(v.Message, InvalidCode) {
			return ierr.InvalidOTP{
				PhoneNumber: phoneNumber,
				OTP:         otp,
			}, true
		}
	}

	return ierr.InvalidOTP{}, false
}
