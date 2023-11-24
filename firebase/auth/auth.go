package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	apiKey string
	url    string
}

func (c *Client) SendVerificationCode(ctx context.Context, body SendVerificationCodeRequest) (SendVerificationCodeResponse, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return SendVerificationCodeResponse{}, fmt.Errorf("failed to marshal json: %w", err)
	}

	url := fmt.Sprintf("%s:%s", c.url, "sendVerificationCode")
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(b))
	if err != nil {
		return SendVerificationCodeResponse{}, fmt.Errorf("failed to create new req with context: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	q := req.URL.Query()
	q.Add("key", c.apiKey)
	req.URL.RawQuery = q.Encode()

	h := &http.Client{Timeout: 10 * time.Second}
	resp, err := h.Do(req)
	if err != nil {
		return SendVerificationCodeResponse{}, fmt.Errorf("failed to do request: %w", err)
	}
	defer resp.Body.Close()

	var respBody SendVerificationCodeResponse
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		return SendVerificationCodeResponse{}, fmt.Errorf("failed to decode response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return SendVerificationCodeResponse{}, fmt.Errorf("failed to call firebase api with error code %d and body %v", resp.StatusCode, respBody.Error)
	}

	return respBody, nil
}

func (c *Client) SignInWithPhoneNumber(ctx context.Context, body SignInWithPhoneNumberRequest) (SignInWithPhoneNumberResponse, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return SignInWithPhoneNumberResponse{}, fmt.Errorf("failed to marshal json: %w", err)
	}

	url := fmt.Sprintf("%s:%s", c.url, "signInWithPhoneNumber")
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(b))
	if err != nil {
		return SignInWithPhoneNumberResponse{}, fmt.Errorf("failed to create new req with context: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	q := req.URL.Query()
	q.Add("key", c.apiKey)
	req.URL.RawQuery = q.Encode()

	h := &http.Client{Timeout: 10 * time.Second}
	resp, err := h.Do(req)
	if err != nil {
		return SignInWithPhoneNumberResponse{}, fmt.Errorf("failed to do request: %w", err)
	}
	defer resp.Body.Close()

	var respBody SignInWithPhoneNumberResponse
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		return SignInWithPhoneNumberResponse{}, fmt.Errorf("failed to decode response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		if invalidCodeErr, ok := getInvalidCodeError(respBody.Error, body.PhoneNumber, body.Code); ok {
			return SignInWithPhoneNumberResponse{}, invalidCodeErr
		}

		return SignInWithPhoneNumberResponse{}, fmt.Errorf("failed to call firebase api with error code %d and body %v", resp.StatusCode, respBody.Error)
	}

	return respBody, nil
}

func NewClient(APIKey, URL string) *Client {
	return &Client{
		apiKey: APIKey,
		url:    URL,
	}
}
