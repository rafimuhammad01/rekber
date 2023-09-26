package transaction

import "github.com/google/uuid"

type Transaction struct {
	ID     uuid.UUID
	Seller Seller
	Buyer  Buyer
}
