package user

import (
	"context"
	"fmt"
	"rekber/internal/dto"
	"time"

	"github.com/google/uuid"
)

type OTPRepository interface {
	VerifyOTP(ctx context.Context, phoneNumber, otp string) error
	SendOTP(ctx context.Context, phoneNumber string) error
	SaveVerifiedOTP(ctx context.Context, phoneNumber string, state int) error
	GetVerifiedOTP(ctx context.Context, phoneNumber string, state int) error
}

type Repository interface {
	GetByPhoneNumber(ctx context.Context, phoneNumber string) (User, error)
	Save(ctx context.Context, u User) error
}

type Service struct {
	otpRepository OTPRepository
	repository    Repository
}

func (s Service) Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error) {
	if err := s.otpRepository.VerifyOTP(ctx, req.PhoneNumber, req.OTP); err != nil {
		return dto.LoginResponse{}, fmt.Errorf("failed to verify otp: %w", err)
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

func (s Service) VerifyOTP(ctx context.Context, req dto.VerifyOTP) error {
	if err := s.otpRepository.VerifyOTP(ctx, req.PhoneNumber, req.OTP); err != nil {
		return fmt.Errorf("failed to verify otp: %w", err)
	}

	if err := s.otpRepository.SaveVerifiedOTP(ctx, req.PhoneNumber, int(req.State)); err != nil {
		return fmt.Errorf("failed to save verified otp: %w", err)
	}

	return nil
}

func (s Service) SendOTP(ctx context.Context, req dto.SendOTPRequest) error {
	return nil
}

func (s Service) Register(ctx context.Context, req dto.RegisterRequest) error {
	if err := s.otpRepository.GetVerifiedOTP(ctx, req.PhoneNumber, int(dto.RegisterState)); err != nil {
		return fmt.Errorf("failed to get verified otp: %w", err)
	}

	user := User{
		ID:                    uuid.New(),
		Name:                  req.Name,
		PhoneNumber:           req.PhoneNumber,
		PhoneNumberVerifiedAt: time.Now(), // will register using OTP means that phone number is also verified
		CreatedAt:             time.Now(),
	}

	if err := s.repository.Save(ctx, user); err != nil {
		return fmt.Errorf("failed to save user: %w", err)
	}

	return nil
}

func NewService(userRepo Repository, otpRepo OTPRepository) *Service {
	return &Service{
		otpRepository: otpRepo,
		repository:    userRepo,
	}
}
