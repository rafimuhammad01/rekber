package transaction

import (
	"math/rand"
	"reflect"
	"testing"
	"time"

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
		// TODO: Add test cases.
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
