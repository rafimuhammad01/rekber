package transaction

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Seller struct {
	ID                    uuid.UUID
	PhoneNumberVerifiedAt time.Time
	BankAccount           BankAccount
}

func (s Seller) IsEligible() bool {
	return (!s.PhoneNumberVerifiedAt.IsZero()) && !(s.BankAccount.ID == uuid.Nil)
}

func (s Seller) Accept(t Transaction) (Transaction, error) {
	if !s.IsEligible() {
		return Transaction{}, errors.New("user is not eligible")
	}

	if t.CreatedBy != buyer {
		return Transaction{}, errors.New("transaction created is not by buyer")
	}

	if isVerified := t.VerifyLastStatus(waitingForPayment); !isVerified {
		return Transaction{}, errors.New("transaction last status is not valid")
	}

	t.Status = waitingForPayment
	t.AcceptedAt = time.Now()
	t.AcceptedBy = seller

	return t, nil
}

func (b Seller) Reject(t Transaction, reason string) (Transaction, error) {
	if isVerified := t.VerifyLastStatus(rejected); !isVerified {
		return Transaction{}, errors.New("transaction last status is not valid")
	}

	if t.CreatedBy != buyer {
		return Transaction{}, errors.New("transaction created not by buyer")
	}

	t.RejectedAt = time.Now()
	t.RejectedBy = seller
	t.RejectedReason = reason
	t.Status = rejected

	return t, nil
}
