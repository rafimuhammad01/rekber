package transaction

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// TODO: need to validate last status first before go to next state
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
		ID:        uuid.New(),
		Seller:    s,
		Buyer:     b,
		CreatedBy: buyer,
		CreatedAt: time.Now(),
		Status:    waitingForApproval,
	}, nil
}

func (b Buyer) Accept(t Transaction) (Transaction, error) {
	if !b.IsEligible() {
		return Transaction{}, errors.New("failed to accept transaction: user is not eligible")
	}

	if t.CreatedBy != seller {
		return Transaction{}, errors.New("failed to accept transaction: transaction should be created by seller to be accepted by buyer")
	}

	t.AcceptedAt = time.Now()
	t.AcceptedBy = buyer

	return t, nil
}

func (b Buyer) Reject(t Transaction, reason string) (Transaction, error) {
	t.RejectedAt = time.Now()
	t.RejectedBy = buyer
	t.RejectedReason = reason

	return t, nil
}
