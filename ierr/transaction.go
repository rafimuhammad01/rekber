package ierr

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type BuyerIsNotEligible struct {
	ID     uuid.UUID
	Reason string
}

func (u BuyerIsNotEligible) Error() string {
	return fmt.Sprintf("user with id %s is not eligible because %s", u.ID.String(), u.Reason)
}

func (u BuyerIsNotEligible) HTTPStatusCode() int {
	return http.StatusBadRequest
}

func (u BuyerIsNotEligible) HTTPMessage() string {
	return u.Error()
}

type TransactionNotCreatedBySeller struct{}

func (u TransactionNotCreatedBySeller) Error() string {
	return "transaction should be created by seller to be accepted by buyer"
}

func (u TransactionNotCreatedBySeller) HTTPStatusCode() int {
	return http.StatusBadRequest
}

func (u TransactionNotCreatedBySeller) HTTPMessage() string {
	return u.Error()
}

type TransactionStatusNotValid struct {
	LastStatus string
	NewStatus  string
}

func (u TransactionStatusNotValid) Error() string {
	return fmt.Sprintf("transaction status %s cannot be updated to %s", u.LastStatus, u.NewStatus)
}

func (u TransactionStatusNotValid) HTTPStatusCode() int {
	return http.StatusBadRequest
}

func (u TransactionStatusNotValid) HTTPMessage() string {
	return u.Error()
}
