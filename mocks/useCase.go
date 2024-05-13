package mocks

import (
	"fase-4-hf-order/internal/core/domain/entity/dto"
	"strings"
)

type MockOrderUseCase struct {
	WantOutNull string
	WantErr     error
}

func (m MockOrderUseCase) UpdateOrderByID(id int64, order dto.RequestOrder) error {
	if m.WantErr != nil && strings.EqualFold("errUpdateOrderByID", m.WantErr.Error()) {
		return m.WantErr
	}
	return nil
}

func (m MockOrderUseCase) GetOrders() error {
	if m.WantErr != nil && strings.EqualFold("errGetOrders", m.WantErr.Error()) {
		return m.WantErr
	}
	return nil
}

func (m MockOrderUseCase) GetOrderByID(id int64) error {
	if m.WantErr != nil && strings.EqualFold("errGetOrderByID", m.WantErr.Error()) {
		return m.WantErr
	}
	return nil
}

func (m MockOrderUseCase) SaveOrder(order dto.RequestOrder) error {
	if m.WantErr != nil && strings.EqualFold("errSaveOrder", m.WantErr.Error()) {
		return m.WantErr
	}
	return nil
}
