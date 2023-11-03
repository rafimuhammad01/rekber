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

func (u User) GenerateToken() (Token, error) {
	accessToken, err := u.generateAccessToken()
	if err != nil {
		return Token{}, fmt.Errorf("failed to generate access token: %w", err)
	}

	refreshToken, err := u.generateRefreshToken()
	if err != nil {
		return Token{}, fmt.Errorf("failed to generate refresh token: %w", err)
	}

	return Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (u User) generateAccessToken() (string, error) {
	expiredAt := jwt.NewNumericDate(time.Now().Add(config.Get().JWT.AccessToken.Duration))
	claims := Claims{
		UserID:           u.ID.String(),
		RegisteredClaims: jwt.RegisteredClaims{Issuer: config.Get().App.Name, ExpiresAt: expiredAt},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(config.Get().JWT.AccessToken.SecretKey))
	if err != nil {
		return "", fmt.Errorf("failed to signed access token: %w ", err)
	}
	return signedToken, nil
}

func (u User) generateRefreshToken() (string, error) {
	expiredAt := jwt.NewNumericDate(time.Now().Add(config.Get().JWT.RefreshToken.Duration))
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{Issuer: config.Get().App.Name, ExpiresAt: expiredAt},
		UserID:           u.ID.String(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(config.Get().JWT.RefreshToken.SecretKey))
	if err != nil {
		return "", fmt.Errorf("failed to signed JWT token: %w ", err)
	}

	return signedToken, nil
}
