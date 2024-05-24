package useCase

import (
	"fase-4-hf-order/internal/core/domain/entity/dto"
	"log"
	"testing"
)

// go test -v -failfast -run ^Test_GetOrderByID$
func Test_GetOrderByID(t *testing.T) {
	type args struct {
		id int64
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				id: 1,
			},
			wantErr: false,
		},
		{
			name: "not_valid_cpf",
			args: args{
				id: 0,
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		uc := NewOrderUseCase()
		t.Run(tc.name, func(*testing.T) {
			err := uc.GetOrderByID(tc.args.id)
			if (!tc.wantErr) && err != nil {
				log.Panicf("unexpected error: %v", err)
			}
		})
	}
}

// go test -v -failfast -run ^Test_SaveOrder$
func Test_SaveOrder(t *testing.T) {

	type args struct {
		reqOrder dto.RequestOrder
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				reqOrder: dto.RequestOrder{
					ID:               0,
					ClientUUID:       "",
					VoucherUUID:      "",
					Status:           "paid",
					VerificationCode: "",
					CreatedAt:        "",
				},
			},
			wantErr: false,
		},
		{
			name: "not_valid",
			args: args{
				reqOrder: dto.RequestOrder{},
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		uc := NewOrderUseCase()
		t.Run(tc.name, func(*testing.T) {
			err := uc.SaveOrder(tc.args.reqOrder)
			if (!tc.wantErr) && err != nil {
				log.Panicf("unexpected error: %v", err)
			}
		})
	}
}

// go test -v -failfast -run ^Test_UpdateOrderByID$
func Test_UpdateOrderByID(t *testing.T) {

	type args struct {
		id       int64
		reqOrder dto.RequestOrder
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				id: 1,
				reqOrder: dto.RequestOrder{
					ID:               0,
					ClientUUID:       "",
					VoucherUUID:      "",
					Status:           "paid",
					VerificationCode: "",
					CreatedAt:        "",
				},
			},
			wantErr: false,
		},
		{
			name: "id_not_valid",
			args: args{
				id:       0,
				reqOrder: dto.RequestOrder{},
			},
			wantErr: true,
		},
		{
			name: "not_valid",
			args: args{
				reqOrder: dto.RequestOrder{},
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		uc := NewOrderUseCase()
		t.Run(tc.name, func(*testing.T) {
			err := uc.UpdateOrderByID(tc.args.id, tc.args.reqOrder)
			if (!tc.wantErr) && err != nil {
				log.Panicf("unexpected error: %v", err)
			}
		})
	}

}
