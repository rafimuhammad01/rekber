package transaction

import (
	"rekber/ierr"
	"time"

	"github.com/google/uuid"
)

type Buyer struct {
	ID                    uuid.UUID
	PhoneNumberVerifiedAt time.Time
}

func (b Buyer) IsEligible() bool {
	return !b.PhoneNumberVerifiedAt.IsZero()
}

func (b Buyer) Create(s Seller) (Transaction, error) {
	if !b.IsEligible() {
		return Transaction{}, ierr.BuyerIsNotEligible{ID: b.ID, Reason: "phone number is not verified yet"}
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
		return Transaction{}, ierr.BuyerIsNotEligible{ID: b.ID, Reason: "phone number is not verified yet"}
	}

	if t.CreatedBy != seller {
		return Transaction{}, ierr.TransactionNotCreatedBySeller{}
	}

	if !t.VerifyLastStatus(waitingForPayment) {
		return Transaction{}, ierr.TransactionStatusNotValid{
			LastStatus: t.Status.String(),
			NewStatus:  waitingForPayment.String(),
		}
	}

	t.AcceptedAt = time.Now()
	t.AcceptedBy = buyer
	t.Status = paid

	return t, nil
}

func (b Buyer) Reject(t Transaction, reason string) (Transaction, error) {
	if !t.VerifyLastStatus(rejected) {
		return Transaction{}, ierr.TransactionStatusNotValid{
			LastStatus: t.Status.String(),
			NewStatus:  rejected.String(),
		}
	}

	if t.CreatedBy != seller {
		return Transaction{}, ierr.TransactionNotCreatedBySeller{}
	}

	t.RejectedAt = time.Now()
	t.RejectedBy = buyer
	t.RejectedReason = reason
	t.Status = rejected

	return t, nil
}

func (b Buyer) Done(t Transaction) (Transaction, error) {
	if !b.IsEligible() {
		return Transaction{}, ierr.BuyerIsNotEligible{ID: b.ID, Reason: "phone number is not verified yet"}
	}

	if !t.VerifyLastStatus(success) {
		return Transaction{}, ierr.TransactionStatusNotValid{
			LastStatus: t.Status.String(),
			NewStatus:  success.String(),
		}
	}

	t.Status = success
	t.SuccessAt = time.Now()

	return t, nil
}
