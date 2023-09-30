package user

import (
	"reflect"
	"rekber/config"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/google/uuid"
)

func TestUser_generateAccessToken(t *testing.T) {
	layoutFormat := "2006-01-02 15:04:05"
	value := "2023-09-30 21:00:00"
	timeNow, _ := time.Parse(layoutFormat, value)
	monkey.Patch(time.Now, func() time.Time {
		return timeNow
	})

	userUUID := uuid.MustParse("e41f16ef-0530-42ed-8b02-4ae2fa4c4dc2")

	type fields struct {
		ID                    uuid.UUID
		PhoneNumber           string
		PhoneNumberVerifiedAt time.Time
		BankAccount           BankAccount
	}
	tests := []struct {
		name    string
		fields  fields
		conf    config.Config
		want    string
		wantErr bool
	}{
		{
			name: "successfully generate access token",
			conf: config.Config{
				AppName:                 "testing-app",
				AccessTokenDuration:     time.Duration(time.Hour),
				JWTAccessTokenSecretKey: "test-secret-key",
			},
			fields: fields{
				ID:                    userUUID,
				PhoneNumber:           "8121313231",
				PhoneNumberVerifiedAt: time.Now(),
			},
			want:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ0ZXN0aW5nLWFwcCIsImV4cCI6MTY5NjExMTIwMCwidXNlcl9pZCI6ImU0MWYxNmVmLTA1MzAtNDJlZC04YjAyLTRhZTJmYTRjNGRjMiJ9.hgxn0WtzqGtZb1mFCjrJ3zqq3Wg_u4fhoziVAo_5hWc",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := User{
				ID:                    tt.fields.ID,
				PhoneNumber:           tt.fields.PhoneNumber,
				PhoneNumberVerifiedAt: tt.fields.PhoneNumberVerifiedAt,
				BankAccount:           tt.fields.BankAccount,
			}

			config.Set(tt.conf)

			got, err := u.generateAccessToken()
			if (err != nil) != tt.wantErr {
				t.Errorf("User.generateAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("User.generateAccessToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_generateRefreshToken(t *testing.T) {
	layoutFormat := "2006-01-02 15:04:05"
	value := "2023-09-30 21:00:00"
	timeNow, _ := time.Parse(layoutFormat, value)
	monkey.Patch(time.Now, func() time.Time {
		return timeNow
	})

	userUUID := uuid.MustParse("e41f16ef-0530-42ed-8b02-4ae2fa4c4dc2")

	type fields struct {
		ID                    uuid.UUID
		PhoneNumber           string
		PhoneNumberVerifiedAt time.Time
		BankAccount           BankAccount
	}
	tests := []struct {
		name    string
		fields  fields
		conf    config.Config
		want    string
		wantErr bool
	}{
		{
			name: "successfully generate refresh token",
			conf: config.Config{
				AppName:                  "testing-app",
				RefreshTokenDuration:     time.Duration(time.Hour),
				JWTRefreshTokenSecretKey: "test-secret-key",
			},
			fields: fields{
				ID:                    userUUID,
				PhoneNumber:           "8121313231",
				PhoneNumberVerifiedAt: time.Now(),
			},
			want:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ0ZXN0aW5nLWFwcCIsImV4cCI6MTY5NjExMTIwMCwidXNlcl9pZCI6ImU0MWYxNmVmLTA1MzAtNDJlZC04YjAyLTRhZTJmYTRjNGRjMiJ9.hgxn0WtzqGtZb1mFCjrJ3zqq3Wg_u4fhoziVAo_5hWc",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := User{
				ID:                    tt.fields.ID,
				PhoneNumber:           tt.fields.PhoneNumber,
				PhoneNumberVerifiedAt: tt.fields.PhoneNumberVerifiedAt,
				BankAccount:           tt.fields.BankAccount,
			}

			config.Set(tt.conf)

			got, err := u.generateRefreshToken()
			if (err != nil) != tt.wantErr {
				t.Errorf("User.generateRefreshToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("User.generateRefreshToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_GenerateToken(t *testing.T) {
	layoutFormat := "2006-01-02 15:04:05"
	value := "2023-09-30 21:00:00"
	timeNow, _ := time.Parse(layoutFormat, value)
	monkey.Patch(time.Now, func() time.Time {
		return timeNow
	})

	config.Set(config.Config{
		AppName:                  "testing-app",
		AccessTokenDuration:      time.Duration(time.Hour),
		JWTAccessTokenSecretKey:  "test-secret-key",
		RefreshTokenDuration:     time.Duration(time.Hour),
		JWTRefreshTokenSecretKey: "test-secret-key",
	})

	userUUID := uuid.MustParse("e41f16ef-0530-42ed-8b02-4ae2fa4c4dc2")

	type fields struct {
		ID                    uuid.UUID
		PhoneNumber           string
		PhoneNumberVerifiedAt time.Time
		BankAccount           BankAccount
	}
	tests := []struct {
		name    string
		fields  fields
		want    Token
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			fields: fields{
				ID:                    userUUID,
				PhoneNumber:           "8121313231",
				PhoneNumberVerifiedAt: time.Now(),
			},
			want: Token{
				AccessToken:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ0ZXN0aW5nLWFwcCIsImV4cCI6MTY5NjExMTIwMCwidXNlcl9pZCI6ImU0MWYxNmVmLTA1MzAtNDJlZC04YjAyLTRhZTJmYTRjNGRjMiJ9.hgxn0WtzqGtZb1mFCjrJ3zqq3Wg_u4fhoziVAo_5hWc",
				RefreshToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ0ZXN0aW5nLWFwcCIsImV4cCI6MTY5NjExMTIwMCwidXNlcl9pZCI6ImU0MWYxNmVmLTA1MzAtNDJlZC04YjAyLTRhZTJmYTRjNGRjMiJ9.hgxn0WtzqGtZb1mFCjrJ3zqq3Wg_u4fhoziVAo_5hWc",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := User{
				ID:                    tt.fields.ID,
				PhoneNumber:           tt.fields.PhoneNumber,
				PhoneNumberVerifiedAt: tt.fields.PhoneNumberVerifiedAt,
				BankAccount:           tt.fields.BankAccount,
			}
			got, err := u.GenerateToken()
			if (err != nil) != tt.wantErr {
				t.Errorf("User.GenerateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.GenerateToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
