package token

import (
	"fmt"
	"rekber/config"
	"rekber/ierr"
	"rekber/internal/user"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Token struct {
	accessToken  string
	refreshToken string
}

func (t Token) ParseAccessToken() (user.User, error) {
	jwtToken, err := validate(t.accessToken, config.Get().JWT.AccessToken.SecretKey)
	if err != nil {
		return user.User{}, fmt.Errorf("failed to validate access token: %w", err)
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !(ok && jwtToken.Valid) {
		return user.User{}, ierr.JWTError{}
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return user.User{}, ierr.JWTError{}
	}

	return user.User{
		ID: uuid.MustParse(userID),
	}, nil
}

func (t Token) ParseRefreshToken() (user.User, error) {
	jwtToken, err := validate(t.refreshToken, config.Get().JWT.RefreshToken.SecretKey)
	if err != nil {
		return user.User{}, fmt.Errorf("failed to validate refresh token: %w", err)
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !(ok && jwtToken.Valid) {
		return user.User{}, ierr.JWTError{}
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return user.User{}, ierr.JWTError{}
	}

	return user.User{
		ID: uuid.MustParse(userID),
	}, nil
}

func validate(token, secretKey string) (*jwt.Token, error) {
	tokenParsed, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ierr.JWTError{}
		}

		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", ierr.JWTError{})
	}

	return tokenParsed, nil
}

type Options func(token *Token)

func NewToken(options ...Options) *Token {
	t := &Token{}

	for _, opt := range options {
		opt(t)
	}

	return t
}

func WithAccessToken(t string) Options {
	return func(token *Token) {
		token.accessToken = t
	}
}
