package transaction

import (
	"math/rand"
	"reflect"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/google/uuid"
)

func TestBuyer_IsEligible(t *testing.T) {
	type fields struct {
		ID                    uuid.UUID
		PhoneNumberVerifiedAt time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "phone number is not verified",
			fields: fields{
				ID: uuid.New(),
			},
			want: false,
		},
		{
			name: "phone number is verified",
			fields: fields{
				ID:                    uuid.New(),
				PhoneNumberVerifiedAt: time.Now(),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Buyer{
				ID:                    tt.fields.ID,
				PhoneNumberVerifiedAt: tt.fields.PhoneNumberVerifiedAt,
			}
			if got := b.IsEligible(); got != tt.want {
				t.Errorf("Buyer.IsEligible() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuyer_Create(t *testing.T) {
	uuid.SetRand(rand.New(rand.NewSource(1)))

	uuidBuyer := uuid.MustParse("861b1cd4-90ec-4633-9e84-dcfbd03a9fe5")
	uuidSeller := uuid.MustParse("28551a5b-c62f-43bb-9893-3438bc6135df")
	verifiedAt := time.Now()

	createdAt := time.Now()
	monkey.Patch(time.Now, func() time.Time {
		return createdAt
	})

	type fields struct {
		ID                    uuid.UUID
		PhoneNumberVerifiedAt time.Time
	}
	type args struct {
		s Seller
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Transaction
		wantErr bool
	}{
		{
			name: "buyer is not eligible",
			fields: fields{
				ID: uuidBuyer,
			},
			args: args{
				s: Seller{
					ID: uuidSeller,
				},
			},
			want:    Transaction{},
			wantErr: true,
		},
		{
			name: "buyer is eligible",
			fields: fields{
				ID:                    uuidBuyer,
				PhoneNumberVerifiedAt: verifiedAt,
			},
			args: args{
				s: Seller{
					ID: uuidSeller,
				},
			},
			want: Transaction{
				ID: uuid.MustParse("52fdfc07-2182-454f-963f-5f0f9a621d72"),
				Seller: Seller{
					ID: uuidSeller,
				},
				Buyer: Buyer{
					ID:                    uuidBuyer,
					PhoneNumberVerifiedAt: verifiedAt,
				},
				CreatedBy: buyer,
				CreatedAt: createdAt,
				Status:    waitingForApproval,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Buyer{
				ID:                    tt.fields.ID,
				PhoneNumberVerifiedAt: tt.fields.PhoneNumberVerifiedAt,
			}
			got, err := b.Create(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Buyer.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Buyer.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuyer_Accept(t *testing.T) {
	trxUUID := uuid.New()

	acceptedAt := time.Now()
	monkey.Patch(time.Now, func() time.Time {
		return acceptedAt
	})

	type fields struct {
		ID                    uuid.UUID
		PhoneNumberVerifiedAt time.Time
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
			name: "buyer accept transaction successfully",
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
			},
			want: Transaction{
				ID:         trxUUID,
				CreatedBy:  seller,
				Status:     paid,
				AcceptedAt: acceptedAt,
				AcceptedBy: buyer,
			},
			wantErr: false,
		},
		{
			name: "buyer is not eligible because doesn't verified yet",
			fields: fields{
				ID: trxUUID,
			},
			args: args{
				t: Transaction{
					ID: trxUUID,
				},
			},
			want:    Transaction{},
			wantErr: true,
		},
		{
			name: "buyer is not eligible because transaction is also created by buyer",
			fields: fields{
				ID:                    trxUUID,
				PhoneNumberVerifiedAt: time.Now(),
			},
			args: args{
				t: Transaction{
					ID:        trxUUID,
					CreatedBy: buyer,
				},
			},
			want:    Transaction{},
			wantErr: true,
		},
		{
			name: "transaction status is not waiting for approval",
			fields: fields{
				ID:                    uuid.New(),
				PhoneNumberVerifiedAt: time.Now(),
			},
			args: args{
				t: Transaction{
					ID:        trxUUID,
					CreatedBy: seller,
					Status:    paid,
				},
			},
			want:    Transaction{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Buyer{
				ID:                    tt.fields.ID,
				PhoneNumberVerifiedAt: tt.fields.PhoneNumberVerifiedAt,
			}
			got, err := b.Accept(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("Buyer.Accept() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Buyer.Accept() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuyer_Reject(t *testing.T) {
	trxUUID := uuid.New()

	rejectedAt := time.Now()
	monkey.Patch(time.Now, func() time.Time {
		return rejectedAt
	})

	type fields struct {
		ID                    uuid.UUID
		PhoneNumberVerifiedAt time.Time
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
			name: "buyer reject transaction successfully",
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
			want: Transaction{
				ID:             trxUUID,
				CreatedBy:      seller,
				RejectedAt:     rejectedAt,
				RejectedBy:     buyer,
				RejectedReason: "test",
				Status:         rejected,
			},
			wantErr: false,
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
					CreatedBy: seller,
					Status:    paid,
				},
				reason: "test",
			},
			want:    Transaction{},
			wantErr: true,
		},
		{
			name: "transaction is not created by seller",
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
			want:    Transaction{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Buyer{
				ID:                    tt.fields.ID,
				PhoneNumberVerifiedAt: tt.fields.PhoneNumberVerifiedAt,
			}
			got, err := b.Reject(tt.args.t, tt.args.reason)
			if (err != nil) != tt.wantErr {
				t.Errorf("Buyer.Accept() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Buyer.Accept() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestBuyer_Done(t *testing.T) {
	trxUUID := uuid.New()

	successAt := time.Now()
	monkey.Patch(time.Now, func() time.Time {
		return successAt
	})

	type fields struct {
		ID                    uuid.UUID
		PhoneNumberVerifiedAt time.Time
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
			name: "buyer set transaction to done successfully",
			fields: fields{
				ID:                    uuid.New(),
				PhoneNumberVerifiedAt: time.Now(),
			},
			args: args{
				t: Transaction{
					ID:     trxUUID,
					Status: doneBySeller,
				},
			},
			want: Transaction{
				ID:        trxUUID,
				Status:    success,
				SuccessAt: successAt,
			},
			wantErr: false,
		},
		{
			name: "buyer is not eligible",
			fields: fields{
				ID: uuid.New(),
			},
			args: args{
				t: Transaction{
					ID:     trxUUID,
					Status: doneBySeller,
				},
			},
			want:    Transaction{},
			wantErr: true,
		},
		{
			name: "transaction last status is not valid",
			fields: fields{
				ID:                    uuid.New(),
				PhoneNumberVerifiedAt: time.Now(),
			},
			args: args{
				t: Transaction{
					ID:     trxUUID,
					Status: success,
				},
			},
			want:    Transaction{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Buyer{
				ID:                    tt.fields.ID,
				PhoneNumberVerifiedAt: tt.fields.PhoneNumberVerifiedAt,
			}
			got, err := b.Done(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("Buyer.Done() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Buyer.Done() = %v, want %v", got, tt.want)
			}
		})
	}
}
