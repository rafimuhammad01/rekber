package auth

type SendVerificationCodeRequest struct {
	PhoneNumber    string `json:"phoneNumber"`
	RecaptchaToken string `json:"recaptchaToken"`
}

type SendVerificationCodeResponse struct {
	SessionInfo string `json:"sessionInfo"`
	Error       Error  `json:"error"`
}

type SignInWithPhoneNumberRequest struct {
	SessionInfo string `json:"sessionInfo"`
	PhoneNumber string `json:"phoneNumber"`
	Code        string `json:"code"`
}

type SignInWithPhoneNumberResponse struct {
	LocalId     string `json:"localId"`
	PhoneNumber string `json:"phoneNumber"`
	Error       Error  `json:"error"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Errors  []struct {
		Message string `json:"message"`
		Domain  string `json:"domain"`
		Reason  string `json:"reason"`
	} `json:"errors"`
}
