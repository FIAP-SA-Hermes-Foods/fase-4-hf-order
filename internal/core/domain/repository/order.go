package repository

import (
	"fase-4-hf-order/internal/core/domain/entity/dto"
)

type OrderRepository interface {
	SaveOrder(order dto.OrderDB) (*dto.OrderDB, error)
	UpdateOrderByID(id int64, order dto.OrderDB) (*dto.OrderDB, error)
	GetOrders() ([]dto.OrderDB, error)
	GetOrderByID(id int64) (*dto.OrderDB, error)
}
