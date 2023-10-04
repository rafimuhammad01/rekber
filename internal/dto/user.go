package dto

const (
	RegisterState VerifyOTPState = iota + 1
)

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	OTP         string `json:"otp"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type SendOTPRequest struct {
	PhoneNumber string `json:"phone_number"`
}

type VerifyOTPState int

type VerifyOTP struct {
	PhoneNumber string         `json:"phone_number"`
	OTP         string         `json:"otp"`
	State       VerifyOTPState `json:"state"`
}

type RegisterRequest struct {
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
}
