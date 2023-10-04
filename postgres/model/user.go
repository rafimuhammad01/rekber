package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID                    uuid.UUID `db:"id"`
	Name                  string    `db:"name"`
	PhoneNumber           string    `db:"phone_number"`
	PhoneNumberVerifiedAt time.Time `db:"phone_number_verified_at"`
	CreatedAt             time.Time `db:"created_at"`
}
