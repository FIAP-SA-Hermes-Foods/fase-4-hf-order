package useCase

import "fase-4-hf-order/internal/core/domain/entity/dto"

type OrderUseCase interface {
	SaveOrder(order dto.RequestOrder) error
	GetOrderByID(id int64) error
	UpdateOrderByID(id int64, order dto.RequestOrder) error
	GetOrders() error
}
