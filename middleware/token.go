package middleware

import (
	"fmt"
	"rekber/config"
	"rekber/ierr"
	"rekber/internal/user"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type token struct {
	accessToken  string
	refreshToken string
}

func (t token) validateAccessToken() (*jwt.Token, error) {
	tokenParsed, err := jwt.Parse(t.accessToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ierr.JWTError{}
		}

		return []byte(config.Get().JWT.AccessToken.SecretKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse refresh token: %w", ierr.JWTError{})
	}

	return tokenParsed, nil
}

func (t token) parseAccessToken() (user.User, error) {
	jwtToken, err := t.validateAccessToken()
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

func (t token) validateRefreshToken() (*jwt.Token, error) {
	tokenParsed, err := jwt.Parse(t.refreshToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ierr.JWTError{}
		}

		return []byte(config.Get().JWT.RefreshToken.SecretKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse refresh token: %w", ierr.JWTError{})
	}

	return tokenParsed, nil
}

func (t token) parseRefreshToken() (user.User, error) {
	jwtToken, err := t.validateRefreshToken()
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
