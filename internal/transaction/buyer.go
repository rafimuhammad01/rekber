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
		return Transaction{}, errors.New("user is not eligible")
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
		return Transaction{}, errors.New("user is not eligible")
	}

	if t.CreatedBy != seller {
		return Transaction{}, errors.New("transaction should be created by seller to be accepted by buyer")
	}

	if isVerified := t.VerifyLastStatus(waitingForPayment); !isVerified {
		return Transaction{}, errors.New("transaction last status is not valid")
	}

	t.AcceptedAt = time.Now()
	t.AcceptedBy = buyer
	t.Status = paid

	return t, nil
}

func (b Buyer) Reject(t Transaction, reason string) (Transaction, error) {
	if isVerified := t.VerifyLastStatus(rejected); !isVerified {
		return Transaction{}, errors.New("transaction last status is not valid")
	}

	if t.CreatedBy != seller {
		return Transaction{}, errors.New("transaction is not created by seller")
	}

	t.RejectedAt = time.Now()
	t.RejectedBy = buyer
	t.RejectedReason = reason
	t.Status = rejected

	return t, nil
}
