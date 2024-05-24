package application

import "fase-4-hf-order/internal/core/domain/entity/dto"

func (app application) UpdateOrderByIDRepository(id int64, order dto.OrderDB) (*dto.OrderDB, error) {
	return app.orderRepo.UpdateOrderByID(id, order)
}

func (app application) GetOrdersRepository() ([]dto.OrderDB, error) {
	return app.orderRepo.GetOrders()
}

func (app application) SaveOrderRepository(order dto.OrderDB) (*dto.OrderDB, error) {
	return app.orderRepo.SaveOrder(order)
}

func (app application) GetOrderByIDRepository(id int64) (*dto.OrderDB, error) {
	return app.orderRepo.GetOrderByID(id)
}

func (app application) GetAllOrderItemRepository() ([]dto.OrderItemDB, error) {
	return app.orderItemRepo.GetAllOrderItem()
}

func (app application) GetAllOrderItemByOrderIDRepository(id int64) ([]dto.OrderItemDB, error) {
	return app.orderItemRepo.GetAllOrderItemByOrderID(id)
}

func (app application) SaveOrderItemRepository(order dto.OrderItemDB) (*dto.OrderItemDB, error) {
	return app.orderItemRepo.SaveOrderItem(order)
}
