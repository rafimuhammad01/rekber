package user

const (
	RegisterState VerifyOTPState = iota + 1
	LoginState
)

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type SendOTPRequest struct {
	PhoneNumber string `json:"phone_number"`
	Captcha     string `json:"captcha"`
}

type SendOTPResponse struct {
	SessionInfo string `json:"session_info"`
}

type VerifyOTPState int

type VerifyOTP struct {
	PhoneNumber string         `json:"phone_number"`
	OTP         string         `json:"otp"`
	State       VerifyOTPState `json:"state"`
	SessionInfo string         `json:"session_info"`
}

type RegisterRequest struct {
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
}
