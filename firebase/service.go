package firebase

import (
	"context"
	"fmt"
	"rekber/firebase/auth"
	"rekber/ierr"
	"sync"
)

type Client struct {
	auth  *auth.Client
	cache map[string]interface{} // TODO: implement with redis
	mu    sync.Mutex
}

func (c *Client) VerifyOTP(ctx context.Context, phoneNumber, otp, sessionInfo string) error {
	_, err := c.auth.SignInWithPhoneNumber(ctx, auth.SignInWithPhoneNumberRequest{
		SessionInfo: sessionInfo,
		PhoneNumber: phoneNumber,
		Code:        otp,
	})
	if err != nil {
		return fmt.Errorf("error when sign in with phone number: %w", err)
	}

	return nil
}

func (c *Client) SendOTP(ctx context.Context, phoneNumber, captcha string) (string, error) {
	sessionInfo, err := c.auth.SendVerificationCode(ctx, auth.SendVerificationCodeRequest{
		PhoneNumber:    phoneNumber,
		RecaptchaToken: captcha,
	})
	if err != nil {
		return "", fmt.Errorf("error when send verifcation code: %w", err)
	}

	return sessionInfo.SessionInfo, nil
}

func (c *Client) SaveVerifiedOTP(ctx context.Context, phoneNumber string, state int) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[phoneNumber] = state
	return nil
}

func (c *Client) GetVerifiedOTP(ctx context.Context, phoneNumber string, state int) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	earlierState, ok := c.cache[phoneNumber].(int)
	if !ok {
		return ierr.UserForbiddenAccess{PhoneNumber: phoneNumber}
	}

	if earlierState != state {
		return ierr.UserForbiddenAccess{PhoneNumber: phoneNumber}
	}

	return nil
}

func NewClient(APIKey string, authURL string) *Client {
	return &Client{
		auth:  auth.NewClient(APIKey, authURL),
		cache: make(map[string]interface{}),
	}
}
