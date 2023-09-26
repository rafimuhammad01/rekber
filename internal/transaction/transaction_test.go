package transaction

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestTransaction_VerifyLastStatus(t *testing.T) {
	type fields struct {
		ID             uuid.UUID
		Seller         Seller
		Buyer          Buyer
		CreatedBy      Actors
		CreatedAt      time.Time
		AcceptedAt     time.Time
		AcceptedBy     Actors
		RejectedAt     time.Time
		RejectedBy     Actors
		RejectedReason string
		PaidAt         time.Time
		Status         Status
	}
	type args struct {
		currentStatus Status
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{
			name:   "",
			fields: fields{},
			args:   args{},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Transaction{
				ID:             tt.fields.ID,
				Seller:         tt.fields.Seller,
				Buyer:          tt.fields.Buyer,
				CreatedBy:      tt.fields.CreatedBy,
				CreatedAt:      tt.fields.CreatedAt,
				AcceptedAt:     tt.fields.AcceptedAt,
				AcceptedBy:     tt.fields.AcceptedBy,
				RejectedAt:     tt.fields.RejectedAt,
				RejectedBy:     tt.fields.RejectedBy,
				RejectedReason: tt.fields.RejectedReason,
				PaidAt:         tt.fields.PaidAt,
				Status:         tt.fields.Status,
			}
			if got := tr.VerifyLastStatus(tt.args.currentStatus); got != tt.want {
				t.Errorf("Transaction.VerifyLastStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}
