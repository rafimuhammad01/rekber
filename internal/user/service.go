package user

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type OTPRepository interface {
	VerifyOTP(ctx context.Context, phoneNumber, otp, sessionInfo string) error
	SendOTP(ctx context.Context, phoneNumber, captcha string) (string, error)
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

func (s Service) Login(ctx context.Context, req LoginRequest) (LoginResponse, error) {
	if err := s.otpRepository.GetVerifiedOTP(ctx, req.PhoneNumber, int(LoginState)); err != nil {
		return LoginResponse{}, fmt.Errorf("failed to verify otp: %w", err)
	}

	user, err := s.repository.GetByPhoneNumber(ctx, req.PhoneNumber)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("failed to get user by phone number: %w", err)
	}

	generatedToken, err := user.generateToken()
	if err != nil {
		return LoginResponse{}, fmt.Errorf("failed to generate token: %w", err)
	}

	return LoginResponse{
		AccessToken:  generatedToken.accessToken,
		RefreshToken: generatedToken.refreshToken,
	}, nil
}

func (s Service) VerifyOTP(ctx context.Context, req VerifyOTP) error {
	if err := s.otpRepository.VerifyOTP(ctx, req.PhoneNumber, req.OTP, req.SessionInfo); err != nil {
		return fmt.Errorf("failed to verify otp: %w", err)
	}

	if err := s.otpRepository.SaveVerifiedOTP(ctx, req.PhoneNumber, int(req.State)); err != nil {
		return fmt.Errorf("failed to save verified otp: %w", err)
	}

	return nil
}

func (s Service) SendOTP(ctx context.Context, req SendOTPRequest) (SendOTPResponse, error) {
	sessionInfo, err := s.otpRepository.SendOTP(ctx, req.PhoneNumber, req.Captcha)
	if err != nil {
		return SendOTPResponse{}, err
	}

	return SendOTPResponse{
		SessionInfo: sessionInfo,
	}, nil
}

func (s Service) Register(ctx context.Context, req RegisterRequest) error {
	if err := s.otpRepository.GetVerifiedOTP(ctx, req.PhoneNumber, int(RegisterState)); err != nil {
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
