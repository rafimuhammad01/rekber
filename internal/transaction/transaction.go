package transaction

import (
	"time"

	"github.com/google/uuid"
)

type Actors int

const (
	buyer Actors = iota + 1
	seller
)

type Status int

const (
	// before transfer
	waitingForApproval Status = iota + 1
	waitingForPayment         // also means accepted
	rejected
	paid
	expired

	// after transfer by buyer
	doneBySeller

	// transfer to seller
	success // also means done by buyer
)

type Transaction struct {
	ID uuid.UUID

	// Actors information
	Seller Seller
	Buyer  Buyer

	// Creation information
	CreatedBy Actors
	CreatedAt time.Time

	// Accepted information
	AcceptedAt time.Time
	AcceptedBy Actors

	// Rejected information
	RejectedAt     time.Time
	RejectedBy     Actors
	RejectedReason string

	// Payment information
	PaidAt time.Time

	// Done information
	SuccessAt      time.Time
	DoneBySellerAt time.Time

	// State information
	Status Status
}

func (t Transaction) VerifyLastStatus(updated Status) bool {
	switch t.Status {
	case waitingForApproval:
		return (updated == waitingForPayment) || (updated == rejected)
	case waitingForPayment:
		return (updated == paid) || (updated == expired)
	case paid:
		return (updated == doneBySeller)
	case doneBySeller:
		return (updated == success)
	default:
		return false
	}
}
