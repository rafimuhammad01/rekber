package user

import (
	"context"
	"fmt"
	"rekber/internal/dto"
)

type OTPRepository interface {
	VerifyOTP(ctx context.Context, phoneNumber, otp string) error
}

type Repository interface {
	GetByPhoneNumber(ctx context.Context, phoneNumber string) (User, error)
}

type Service struct {
	otpRepository OTPRepository
	repository    Repository
}

func (s Service) Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error) {
	if err := s.otpRepository.VerifyOTP(ctx, req.PhoneNumber, req.OTP); err != nil {
		return dto.LoginResponse{}, fmt.Errorf("failed to verify otp: %w ", err)
	}

	user, err := s.repository.GetByPhoneNumber(ctx, req.PhoneNumber)
	if err != nil {
		return dto.LoginResponse{}, fmt.Errorf("failed to get user by phone number: %w", err)
	}

	token, err := user.GenerateToken()
	if err != nil {
		return dto.LoginResponse{}, fmt.Errorf("failed to generate token: %w", err)
	}

	return dto.LoginResponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, nil
}
