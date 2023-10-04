package otp

import (
	"context"
)

type Repository struct{}

func (r Repository) VerifyOTP(ctx context.Context, phoneNumber, otp string) error {
	return nil
}

func (r Repository) SendOTP(ctx context.Context, phoneNumber string) error {
	return nil
}

func (r Repository) SaveVerifiedOTP(ctx context.Context, phoneNumber string, state int) error {
	return nil
}

func (r Repository) GetVerifiedOTP(ctx context.Context, phoneNumber string, state int) error {
	return nil
}

func NewRepository() *Repository {
	return &Repository{}
}
