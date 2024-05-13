package dto

import (
	"fase-4-hf-order/internal/core/domain/entity"
	vo "fase-4-hf-order/internal/core/domain/entity/valueObject"
)

type OrderDB struct {
	ID               int64               `json:"id,omitempty"`
	ClientUUID       string              `json:"clientUuid,omitempty"`
	VoucherUUID      string              `json:"voucherUuid,omitempty"`
	Items            []entity.OrderItems `json:"items,omitempty"`
	Status           string              `json:"status,omitempty"`
	VerificationCode string              `json:"verificationCode,omitempty"`
	CreatedAt        string              `json:"createdAt,omitempty"`
}

type (
	RequestOrder struct {
		ID               int64               `json:"id,omitempty"`
		ClientUUID       string              `json:"clientUuid,omitempty"`
		VoucherUUID      string              `json:"voucherUuid,omitempty"`
		Items            []entity.OrderItems `json:"items,omitempty"`
		Status           string              `json:"status,omitempty"`
		VerificationCode string              `json:"verificationCode,omitempty"`
		CreatedAt        string              `json:"createdAt,omitempty"`
	}

	OutputOrder struct {
		ID               int64               `json:"id,omitempty"`
		ClientUUID       string              `json:"clientUuid,omitempty"`
		VoucherUUID      string              `json:"voucherUuid,omitempty"`
		Items            []entity.OrderItems `json:"items,omitempty"`
		Status           string              `json:"status,omitempty"`
		VerificationCode string              `json:"verificationCode,omitempty"`
		CreatedAt        string              `json:"createdAt,omitempty"`
	}
)

func (r RequestOrder) Order() entity.Order {
	return entity.Order{
		ClientUUID:  r.ClientUUID,
		VoucherUUID: r.VoucherUUID,
		Items:       r.Items,
		Status: vo.Status{
			Value: r.Status,
		},
	}
}
