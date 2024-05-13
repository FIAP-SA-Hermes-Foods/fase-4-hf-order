package application

import "fase-4-hf-order/internal/core/domain/entity/dto"

func (app application) UpdateOrderByIDUseCase(id int64, order dto.RequestOrder) error {
	return app.orderUC.UpdateOrderByID(id, order)
}

func (app application) GetOrdersUseCase() error {
	return app.orderUC.GetOrders()
}

func (app application) GetOrderByIDUseCase(id int64) error {
	return app.orderUC.GetOrderByID(id)
}

func (app application) SaveOrderUseCase(order dto.RequestOrder) error {
	return app.orderUC.SaveOrder(order)
}
