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

	// after transfer by buyer
	doneBySeller
	doneByBuyer

	// transfer to seller
	success
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

	// State information
	Status Status
}

// TODO: implement me
func (t Transaction) VerifyLastStatus(currentStatus Status) bool {
	return false
}
