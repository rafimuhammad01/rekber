package transaction

import (
	"reflect"
	"testing"
	"time"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/google/uuid"
)

func TestSeller_IsEligible(t *testing.T) {
	type fields struct {
		ID                    uuid.UUID
		PhoneNumberVerifiedAt time.Time
		BankAccount           BankAccount
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "seller is eligible",
			fields: fields{
				ID:                    uuid.New(),
				PhoneNumberVerifiedAt: time.Now(),
				BankAccount: BankAccount{
					ID: uuid.New(),
				},
			},
			want: true,
		},
		{
			name: "seller is not eligible because phone number is not verified yet",
			fields: fields{
				ID: uuid.New(),
				BankAccount: BankAccount{
					ID: uuid.New(),
				},
			},
			want: false,
		},
		{
			name: "seller is not eligible because bank account is not verified",
			fields: fields{
				ID:                    uuid.New(),
				PhoneNumberVerifiedAt: time.Now(),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Seller{
				ID:                    tt.fields.ID,
				PhoneNumberVerifiedAt: tt.fields.PhoneNumberVerifiedAt,
				BankAccount:           tt.fields.BankAccount,
			}
			if got := s.IsEligible(); got != tt.want {
				t.Errorf("Seller.IsEligible() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSeller_Accept(t *testing.T) {
	trxUUID := uuid.New()
	createdAt := time.Now()

	acceptedAt := time.Now()

	gomonkey.ApplyFunc(time.Now, func() time.Time {
		return acceptedAt
	})

	type fields struct {
		ID                    uuid.UUID
		PhoneNumberVerifiedAt time.Time
		BankAccount           BankAccount
	}
	type args struct {
		t Transaction
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Transaction
		wantErr bool
	}{
		{
			name: "seller success accept transaction",
			fields: fields{
				ID:                    uuid.New(),
				PhoneNumberVerifiedAt: time.Now(),
				BankAccount: BankAccount{
					ID: uuid.New(),
				},
			},
			args: args{
				t: Transaction{
					ID:        trxUUID,
					CreatedBy: buyer,
					CreatedAt: createdAt,
					Status:    waitingForApproval,
				},
			},
			want: Transaction{
				ID:         trxUUID,
				Status:     waitingForPayment,
				AcceptedAt: acceptedAt,
				AcceptedBy: seller,
				CreatedBy:  buyer,
				CreatedAt:  createdAt,
			},
			wantErr: false,
		},
		{
			name: "seller is not eligible",
			fields: fields{
				ID: trxUUID,
			},
			args: args{
				t: Transaction{
					ID:        trxUUID,
					CreatedBy: buyer,
					CreatedAt: createdAt,
				},
			},
			want:    Transaction{},
			wantErr: true,
		},
		{
			name: "transaction not created by buyer",
			fields: fields{
				ID:                    uuid.New(),
				PhoneNumberVerifiedAt: time.Now(),
				BankAccount: BankAccount{
					ID: uuid.New(),
				},
			},
			args: args{
				t: Transaction{
					ID:        trxUUID,
					CreatedBy: seller,
					CreatedAt: createdAt,
				},
			},
			want:    Transaction{},
			wantErr: true,
		},
		{
			name: "last transaction status is not waiting for approval",
			fields: fields{
				ID:                    uuid.New(),
				PhoneNumberVerifiedAt: time.Now(),
				BankAccount: BankAccount{
					ID: uuid.New(),
				},
			},
			args: args{
				t: Transaction{
					ID:        trxUUID,
					CreatedBy: buyer,
					CreatedAt: createdAt,
					Status:    paid,
				},
			},
			want:    Transaction{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Seller{
				ID:                    tt.fields.ID,
				PhoneNumberVerifiedAt: tt.fields.PhoneNumberVerifiedAt,
				BankAccount:           tt.fields.BankAccount,
			}
			got, err := s.Accept(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("Seller.Accept() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Seller.Accept() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestSeller_Reject(t *testing.T) {
	trxUUID := uuid.New()

	rejectedAt := time.Now()
	gomonkey.ApplyFunc(time.Now, func() time.Time {
		return rejectedAt
	})

	type fields struct {
		ID                    uuid.UUID
		PhoneNumberVerifiedAt time.Time
		BankAccount           BankAccount
	}
	type args struct {
		t      Transaction
		reason string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Transaction
		wantErr bool
	}{
		{
			name: "seller reject transaction successfully",
			fields: fields{
				ID:                    uuid.New(),
				PhoneNumberVerifiedAt: time.Now(),
			},
			args: args{
				t: Transaction{
					ID:        trxUUID,
					CreatedBy: buyer,
					Status:    waitingForApproval,
				},
				reason: "test",
			},
			want: Transaction{
				ID:             trxUUID,
				CreatedBy:      buyer,
				RejectedAt:     rejectedAt,
				RejectedBy:     seller,
				RejectedReason: "test",
				Status:         rejected,
			},
			wantErr: false,
		},
		{
			name: "transaction is not created by buyer",
			fields: fields{
				ID:                    uuid.New(),
				PhoneNumberVerifiedAt: time.Now(),
			},
			args: args{
				t: Transaction{
					ID:        trxUUID,
					CreatedBy: seller,
					Status:    waitingForApproval,
				},
				reason: "test",
			},
			want:    Transaction{},
			wantErr: true,
		},
		{
			name: "transaction status is not valid",
			fields: fields{
				ID:                    trxUUID,
				PhoneNumberVerifiedAt: time.Now(),
			},
			args: args{
				t: Transaction{
					ID:        trxUUID,
					CreatedBy: buyer,
					Status:    paid,
				},
				reason: "test",
			},
			want:    Transaction{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Seller{
				ID:                    tt.fields.ID,
				PhoneNumberVerifiedAt: tt.fields.PhoneNumberVerifiedAt,
				BankAccount:           tt.fields.BankAccount,
			}
			got, err := b.Reject(tt.args.t, tt.args.reason)
			if (err != nil) != tt.wantErr {
				t.Errorf("Seller.Reject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Seller.Reject() = %v, want %v", got, tt.want)
			}
		})
	}
}
