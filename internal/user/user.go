package user

import (
	"time"

	"github.com/google/uuid"
)

const (
	buyer Actors = iota + 1
	seller
)

type Actors int

type Bank struct {
	Code string
	Name string
}

type BankAccount struct {
	ID     uuid.UUID
	Number string
	Name   string
	Bank   Bank
}

type User struct {
	ID                    uuid.UUID
	Type                  Actors
	PhoneNumber           string
	PhoneNumberVerifiedAt time.Time
	BankAccount           BankAccount
}
