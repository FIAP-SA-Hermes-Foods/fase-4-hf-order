package mocks

import (
	"fase-4-hf-order/internal/core/domain/entity/dto"
	"strings"
)

type MockOrderRepository struct {
	WantOut     *dto.OrderDB
	WantOutList []dto.OrderDB
	WantErr     error
	WantOutNull string
}

func (m MockOrderRepository) UpdateOrderByID(id int64, order dto.OrderDB) (*dto.OrderDB, error) {
	if m.WantErr != nil && strings.EqualFold("errUpdateOrderByID", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	if strings.EqualFold(m.WantOutNull, "nilUpdateOrderByID") {
		return nil, nil
	}
	return m.WantOut, nil
}

func (m MockOrderRepository) GetOrders() ([]dto.OrderDB, error) {
	if m.WantErr != nil && strings.EqualFold("errGetOrders", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	if strings.EqualFold(m.WantOutNull, "nilGetOrders") {
		return nil, nil
	}
	return m.WantOutList, nil
}

func (m MockOrderRepository) GetOrderByID(id int64) (*dto.OrderDB, error) {
	if m.WantErr != nil && strings.EqualFold("errGetOrderByID", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	if strings.EqualFold(m.WantOutNull, "nilGetOrderByID") {
		return nil, nil
	}
	return m.WantOut, nil
}

func (m MockOrderRepository) SaveOrder(order dto.OrderDB) (*dto.OrderDB, error) {
	if m.WantErr != nil && strings.EqualFold("errSaveOrder", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	if strings.EqualFold(m.WantOutNull, "nilSaveOrder") {
		return nil, nil
	}
	return m.WantOut, nil
}

type MockOrderItemRepository struct {
	WantOut     *dto.OrderItemDB
	WantOutList []dto.OrderItemDB
	WantErr     error
	WantOutNull string
}

func (m MockOrderItemRepository) GetAllOrderItem() ([]dto.OrderItemDB, error) {
	if m.WantErr != nil && strings.EqualFold("errGetAllOrderItem", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	if strings.EqualFold(m.WantOutNull, "nilGetAllOrderItem") {
		return nil, nil
	}
	return m.WantOutList, nil
}

func (m MockOrderItemRepository) GetAllOrderItemByOrderID(id int64) ([]dto.OrderItemDB, error) {
	if m.WantErr != nil && strings.EqualFold("errGetAllOrderItemByOrderID", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	if strings.EqualFold(m.WantOutNull, "nilGetAllOrderItemByOrderID") {
		return nil, nil
	}
	return m.WantOutList, nil
}

func (m MockOrderItemRepository) SaveOrderItem(order dto.OrderItemDB) (*dto.OrderItemDB, error) {
	if m.WantErr != nil && strings.EqualFold("errSaveOrderItem", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	if strings.EqualFold(m.WantOutNull, "nilSaveOrderItem") {
		return nil, nil
	}
	return m.WantOut, nil
}
