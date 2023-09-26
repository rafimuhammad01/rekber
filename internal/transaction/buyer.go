package transaction

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Buyer struct {
	ID                    uuid.UUID
	PhoneNumberVerifiedAt time.Time
}

func (b Buyer) IsEligible() bool {
	return (!b.PhoneNumberVerifiedAt.IsZero())
}

func (b Buyer) Create(s Seller) (Transaction, error) {
	if !b.IsEligible() {
		return Transaction{}, errors.New("failed to create transaction: user is not eligible")
	}

	return Transaction{
		ID:     uuid.New(),
		Seller: s,
		Buyer:  b,
	}, nil
}
