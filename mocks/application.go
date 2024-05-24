package mocks

import (
	"fase-4-hf-order/internal/core/domain/entity/dto"
	"strings"
)

type MockApplication struct {
	WantOut     *dto.OutputOrder
	WantOutList []dto.OutputOrder
	WantErr     error
	WantOutNull string
}

func (m MockApplication) UpdateOrderByID(id int64, order dto.RequestOrder) (*dto.OutputOrder, error) {
	if m.WantErr != nil && strings.EqualFold("errUpdateOrderByID", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	if strings.EqualFold(m.WantOutNull, "nilUpdateOrderByID") {
		return nil, nil
	}
	return m.WantOut, nil
}

func (m MockApplication) GetOrders() ([]dto.OutputOrder, error) {
	if m.WantErr != nil && strings.EqualFold("errGetOrders", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	if strings.EqualFold(m.WantOutNull, "nilGetOrders") {
		return nil, nil
	}
	return m.WantOutList, nil

}

func (m MockApplication) GetOrderByID(id int64) (*dto.OutputOrder, error) {
	if m.WantErr != nil && strings.EqualFold("errGetOrderByID", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	if strings.EqualFold(m.WantOutNull, "nilGetOrderByID") {
		return nil, nil
	}
	return m.WantOut, nil

}

func (m MockApplication) SaveOrder(order dto.RequestOrder) (*dto.OutputOrder, error) {
	if m.WantErr != nil && strings.EqualFold("errSaveOrder", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	if strings.EqualFold(m.WantOutNull, "nilSaveOrder") {
		return nil, nil
	}
	return m.WantOut, nil

}

// Repository Callers
type MockApplicationRepostoryCallers struct {
	WantOut     *dto.OrderDB
	WantOutList []dto.OrderDB
	WantErr     error
}

func (m MockApplicationRepostoryCallers) UpdateOrderByIDRepository(id int64, order dto.OrderDB) (*dto.OrderDB, error) {
	if m.WantErr != nil && strings.EqualFold("errUpdateOrderByIDRepository", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	return m.WantOut, nil
}

func (m MockApplicationRepostoryCallers) GetOrdersRepository() ([]dto.OrderDB, error) {
	if m.WantErr != nil && strings.EqualFold("errGetOrdersRepository", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	return m.WantOutList, nil
}

func (m MockApplicationRepostoryCallers) GetOrderByIDRepository(id int64) (*dto.OrderDB, error) {
	if m.WantErr != nil && strings.EqualFold("errGetOrderByIDRepository", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	return m.WantOut, nil
}

func (m MockApplicationRepostoryCallers) SaveOrderRepository(order dto.OrderDB) (*dto.OrderDB, error) {
	if m.WantErr != nil && strings.EqualFold("errSaveOrderRepository", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	return m.WantOut, nil
}

// UseCase callers
type MockApplicationUseCaseCallers struct {
	WantErr error
}

func (m MockApplicationUseCaseCallers) UpdateOrderByIDUseCase(id int64, order dto.RequestOrder) error {
	if m.WantErr != nil && strings.EqualFold("errUpdateOrderByIDUseCase", m.WantErr.Error()) {
		return m.WantErr
	}
	return nil
}

func (m MockApplicationUseCaseCallers) GetOrdersUseCase() error {
	if m.WantErr != nil && strings.EqualFold("errGetOrdersUseCase", m.WantErr.Error()) {
		return m.WantErr
	}
	return nil
}

func (m MockApplicationUseCaseCallers) GetOrderByIDUseCase(id int64) error {
	if m.WantErr != nil && strings.EqualFold("errGetOrderByIDUseCase", m.WantErr.Error()) {
		return m.WantErr
	}
	return nil
}

func (m MockApplicationUseCaseCallers) SaveOrderUseCase(order dto.RequestOrder) error {
	if m.WantErr != nil && strings.EqualFold("errSaveOrderUseCase", m.WantErr.Error()) {
		return m.WantErr
	}
	return nil
}
