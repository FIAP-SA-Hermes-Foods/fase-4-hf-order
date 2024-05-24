package repository

import "fase-4-hf-order/internal/core/domain/entity/dto"

type OrderItemRepository interface {
	GetAllOrderItem() ([]dto.OrderItemDB, error)
	GetAllOrderItemByOrderID(id int64) ([]dto.OrderItemDB, error)
	SaveOrderItem(order dto.OrderItemDB) (*dto.OrderItemDB, error)
}
