package middleware

import (
	"encoding/base64"
	"reflect"
	"rekber/config"
	"rekber/internal/user"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func TestToken_validateAccessToken(t *testing.T) {
	config.Set(config.Config{
		JWT: config.JWTConfig{
			AccessToken: config.TokenConfig{
				Duration:  time.Hour,
				SecretKey: "test-secret-key",
			},
		},
	})
	secretDecoded, _ := base64.RawURLEncoding.DecodeString("vtV-qYPVoJ8dCEoR98aBr5XsA0gsJwkvQdRpzY7UQQU")

	type fields struct {
		AccessToken  string
		RefreshToken string
	}

	tests := []struct {
		name    string
		fields  fields
		want    *jwt.Token
		wantErr bool
	}{
		{
			name: "success validate access token",
			fields: fields{
				AccessToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ0ZXN0aW5nLWFwcCIsImV4cCI6MjU0ODYyNzIwMCwidXNlcl9pZCI6ImU0MWYxNmVmLTA1MzAtNDJlZC04YjAyLTRhZTJmYTRjNGRjMiJ9.vtV-qYPVoJ8dCEoR98aBr5XsA0gsJwkvQdRpzY7UQQU",
			},
			want: &jwt.Token{
				Raw:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ0ZXN0aW5nLWFwcCIsImV4cCI6MjU0ODYyNzIwMCwidXNlcl9pZCI6ImU0MWYxNmVmLTA1MzAtNDJlZC04YjAyLTRhZTJmYTRjNGRjMiJ9.vtV-qYPVoJ8dCEoR98aBr5XsA0gsJwkvQdRpzY7UQQU",
				Method: jwt.SigningMethodHS256,
				Header: map[string]interface{}{
					"alg": "HS256",
					"typ": "JWT",
				},
				Claims: jwt.MapClaims{
					"exp":     2.5486272e+09,
					"user_id": "e41f16ef-0530-42ed-8b02-4ae2fa4c4dc2",
					"iss":     "testing-app",
				},
				Signature: secretDecoded,
				Valid:     true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := token{
				accessToken:  tt.fields.AccessToken,
				refreshToken: tt.fields.RefreshToken,
			}

			got, err := tr.validateAccessToken()
			if (err != nil) != tt.wantErr {
				t.Errorf("token.validateAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("token.validateAccessToken() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestToken_ParseAccessToken(t *testing.T) {
	config.Set(config.Config{
		JWT: config.JWTConfig{
			AccessToken: config.TokenConfig{
				Duration:  time.Hour,
				SecretKey: "test-secret-key",
			},
		},
	})

	type fields struct {
		AccessToken  string
		RefreshToken string
	}
	tests := []struct {
		name    string
		fields  fields
		want    user.User
		wantErr bool
	}{
		{
			name: "success to parse access token",
			fields: fields{
				AccessToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ0ZXN0aW5nLWFwcCIsImV4cCI6MjU0ODYyNzIwMCwidXNlcl9pZCI6ImU0MWYxNmVmLTA1MzAtNDJlZC04YjAyLTRhZTJmYTRjNGRjMiJ9.vtV-qYPVoJ8dCEoR98aBr5XsA0gsJwkvQdRpzY7UQQU",
			},
			want: user.User{
				ID: uuid.MustParse("e41f16ef-0530-42ed-8b02-4ae2fa4c4dc2"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := token{
				accessToken:  tt.fields.AccessToken,
				refreshToken: tt.fields.RefreshToken,
			}
			got, err := tr.parseAccessToken()
			if (err != nil) != tt.wantErr {
				t.Errorf("token.parseAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("token.parseAccessToken() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestToken_validateRefreshToken(t *testing.T) {
	config.Set(config.Config{
		JWT: config.JWTConfig{
			RefreshToken: config.TokenConfig{
				Duration:  time.Hour,
				SecretKey: "test-secret-key",
			},
		},
	})
	secretDecoded, _ := base64.RawURLEncoding.DecodeString("vtV-qYPVoJ8dCEoR98aBr5XsA0gsJwkvQdRpzY7UQQU")

	type fields struct {
		AccessToken  string
		RefreshToken string
	}
	tests := []struct {
		name    string
		fields  fields
		want    *jwt.Token
		wantErr bool
	}{
		{
			name: "success validate refresh token",
			fields: fields{
				RefreshToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ0ZXN0aW5nLWFwcCIsImV4cCI6MjU0ODYyNzIwMCwidXNlcl9pZCI6ImU0MWYxNmVmLTA1MzAtNDJlZC04YjAyLTRhZTJmYTRjNGRjMiJ9.vtV-qYPVoJ8dCEoR98aBr5XsA0gsJwkvQdRpzY7UQQU",
			},
			want: &jwt.Token{
				Raw:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ0ZXN0aW5nLWFwcCIsImV4cCI6MjU0ODYyNzIwMCwidXNlcl9pZCI6ImU0MWYxNmVmLTA1MzAtNDJlZC04YjAyLTRhZTJmYTRjNGRjMiJ9.vtV-qYPVoJ8dCEoR98aBr5XsA0gsJwkvQdRpzY7UQQU",
				Method: jwt.SigningMethodHS256,
				Header: map[string]interface{}{
					"alg": "HS256",
					"typ": "JWT",
				},
				Claims: jwt.MapClaims{
					"exp":     2.5486272e+09,
					"user_id": "e41f16ef-0530-42ed-8b02-4ae2fa4c4dc2",
					"iss":     "testing-app",
				},
				Signature: secretDecoded,
				Valid:     true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := token{
				accessToken:  tt.fields.AccessToken,
				refreshToken: tt.fields.RefreshToken,
			}
			got, err := tr.validateRefreshToken()
			if (err != nil) != tt.wantErr {
				t.Errorf("token.validateRefreshToken() error = %#v, wantErr %#v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("token.validateRefreshToken() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestToken_ParseRefreshToken(t *testing.T) {
	config.Set(config.Config{
		JWT: config.JWTConfig{
			RefreshToken: config.TokenConfig{
				Duration:  time.Hour,
				SecretKey: "test-secret-key",
			},
		},
	})

	type fields struct {
		AccessToken  string
		RefreshToken string
	}
	tests := []struct {
		name    string
		fields  fields
		want    user.User
		wantErr bool
	}{
		{
			name: "success to parse access token",
			fields: fields{
				RefreshToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ0ZXN0aW5nLWFwcCIsImV4cCI6MjU0ODYyNzIwMCwidXNlcl9pZCI6ImU0MWYxNmVmLTA1MzAtNDJlZC04YjAyLTRhZTJmYTRjNGRjMiJ9.vtV-qYPVoJ8dCEoR98aBr5XsA0gsJwkvQdRpzY7UQQU",
			},
			want: user.User{
				ID: uuid.MustParse("e41f16ef-0530-42ed-8b02-4ae2fa4c4dc2"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := token{
				accessToken:  tt.fields.AccessToken,
				refreshToken: tt.fields.RefreshToken,
			}
			got, err := tr.parseRefreshToken()
			if (err != nil) != tt.wantErr {
				t.Errorf("token.parseRefreshToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("token.parseRefreshToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
