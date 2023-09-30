package user

import (
	"fmt"
	"rekber/config"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Token struct {
	AccessToken  string
	RefreshToken string
}

type Claims struct {
	jwt.RegisteredClaims
	UserID string `json:"user_id"`
}

func (t Token) validateAccessToken() (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(t.AccessToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.Get().JWTAccessTokenSecretKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse jwt: %w", err)
	}

	return token, nil
}

func (t Token) ParseAccessToken() (User, error) {
	jwtToken, err := t.validateAccessToken()
	if err != nil {
		return User{}, fmt.Errorf("failed to validate access token: %w", err)
	}

	claims, ok := jwtToken.Claims.(*Claims)
	if !(ok && jwtToken.Valid) {
		return User{}, fmt.Errorf("token is expired or invalid")
	}

	return User{
		ID: uuid.MustParse(claims.UserID),
	}, nil
}

func (t Token) validateRefreshToken() (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(t.RefreshToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.Get().JWTRefreshTokenSecretKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse refresh token: %w", err)
	}

	return token, nil
}

func (t Token) ParseRefreshToken() (User, error) {
	jwtToken, err := t.validateRefreshToken()
	if err != nil {
		return User{}, fmt.Errorf("failed to validate refresh token: %w", err)
	}

	claims, ok := jwtToken.Claims.(*Claims)
	if !(ok && jwtToken.Valid) {
		return User{}, fmt.Errorf("token is expired or invalid")
	}

	return User{
		ID: uuid.MustParse(claims.UserID),
	}, nil
}