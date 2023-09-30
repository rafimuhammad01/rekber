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
		update Status
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "status is waiting for approval, next to waiting for payment",
			fields: fields{
				Status: waitingForApproval,
			},
			args: args{
				update: waitingForPayment,
			},
			want: true,
		},
		{
			name: "status is waiting for approval, next to rejected",
			fields: fields{
				Status: waitingForApproval,
			},
			args: args{
				update: rejected,
			},
			want: true,
		},
		{
			name: "status is waiting for approval, next to paid",
			fields: fields{
				Status: waitingForApproval,
			},
			args: args{
				update: paid,
			},
			want: false,
		},
		{
			name: "status is waiting for payment, next to paid",
			fields: fields{
				Status: waitingForPayment,
			},
			args: args{
				update: paid,
			},
			want: true,
		},
		{
			name: "status is waiting for payment, next to expired",
			fields: fields{
				Status: waitingForPayment,
			},
			args: args{
				update: expired,
			},
			want: true,
		},
		{
			name: "status is waiting for payment, next to done by seller",
			fields: fields{
				Status: waitingForPayment,
			},
			args: args{
				update: doneBySeller,
			},
			want: false,
		},
		{
			name: "status is paid, next to done by seller",
			fields: fields{
				Status: paid,
			},
			args: args{
				update: doneBySeller,
			},
			want: true,
		},
		{
			name: "status is paid, next to success",
			fields: fields{
				Status: paid,
			},
			args: args{
				update: success,
			},
			want: false,
		},
		{
			name: "status is done by seller, next to success",
			fields: fields{
				Status: doneBySeller,
			},
			args: args{
				update: success,
			},
			want: true,
		},
		{
			name: "status is done by seller, next to paid",
			fields: fields{
				Status: doneBySeller,
			},
			args: args{
				update: paid,
			},
			want: false,
		},
		{
			name: "status unknown",
			fields: fields{
				Status: -1,
			},
			args: args{
				update: -1,
			},
			want: false,
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

			if got := tr.VerifyLastStatus(tt.args.update); got != tt.want {
				t.Errorf("Transaction.VerifyLastStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}
