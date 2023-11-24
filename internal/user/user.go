package user

import (
	"fmt"
	"rekber/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Bank struct {
	Code string
	Name string
}

type BankAccount struct {
	ID     uuid.UUID
	Number string
	Name   string
	Bank   Bank
}

type User struct {
	ID                    uuid.UUID
	PhoneNumber           string
	Name                  string
	PhoneNumberVerifiedAt time.Time
	BankAccount           BankAccount
	CreatedAt             time.Time
}

type token struct {
	accessToken  string
	refreshToken string
}

func (u User) generateToken() (token, error) {
	accessToken, err := u.generateAccessToken()
	if err != nil {
		return token{}, fmt.Errorf("failed to generate access token: %w", err)
	}

	refreshToken, err := u.generateRefreshToken()
	if err != nil {
		return token{}, fmt.Errorf("failed to generate refresh token: %w", err)
	}

	return token{
		accessToken:  accessToken,
		refreshToken: refreshToken,
	}, nil
}

func (u User) generateAccessToken() (string, error) {
	expiredAt := jwt.NewNumericDate(time.Now().Add(config.Get().JWT.AccessToken.Duration))
	claims := jwt.MapClaims{
		"user_id": u.ID.String(),
		"exp":     expiredAt,
		"iss":     config.Get().App.Name,
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := t.SignedString([]byte(config.Get().JWT.AccessToken.SecretKey))
	if err != nil {
		return "", fmt.Errorf("failed to signed access token: %w ", err)
	}
	return signedToken, nil
}

func (u User) generateRefreshToken() (string, error) {
	expiredAt := jwt.NewNumericDate(time.Now().Add(config.Get().JWT.RefreshToken.Duration))
	claims := jwt.MapClaims{
		"user_id": u.ID.String(),
		"exp":     expiredAt,
		"iss":     config.Get().App.Name,
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := t.SignedString([]byte(config.Get().JWT.RefreshToken.SecretKey))
	if err != nil {
		return "", fmt.Errorf("failed to signed JWT token: %w ", err)
	}

	return signedToken, nil
}
