package application

import (
	"errors"
	l "fase-4-hf-order/external/logger"
	ps "fase-4-hf-order/external/strings"
	"fase-4-hf-order/internal/core/domain/entity/dto"
	vo "fase-4-hf-order/internal/core/domain/entity/valueObject"
	"fase-4-hf-order/internal/core/domain/repository"
	"fase-4-hf-order/internal/core/domain/useCase"
	"fmt"
	"strings"
	"time"
)

type Application interface {
	UpdateOrderByID(id int64, order dto.RequestOrder) (*dto.OutputOrder, error)
	GetOrders() ([]dto.OutputOrder, error)
	GetOrderByID(id int64) (*dto.OutputOrder, error)
	SaveOrder(order dto.RequestOrder) (*dto.OutputOrder, error)
}

type application struct {
	orderRepo     repository.OrderRepository
	orderItemRepo repository.OrderItemRepository
	orderUC       useCase.OrderUseCase
}

func NewApplication(orderRepo repository.OrderRepository, orderItemRepo repository.OrderItemRepository, orderUC useCase.OrderUseCase) Application {
	return application{orderRepo: orderRepo, orderItemRepo: orderItemRepo, orderUC: orderUC}
}

func (app application) UpdateOrderByID(id int64, reqOrder dto.RequestOrder) (*dto.OutputOrder, error) {
	l.Infof("UpdateOrderByIDApp: ", " | ", id, " | ", ps.MarshalString(reqOrder))

	foundOrder, err := app.GetOrderByID(id)

	if err != nil {
		l.Errorf("UpdateOrderByIDApp error: ", " | ", err)
		return nil, err
	}

	if foundOrder == nil {
		l.Errorf("UpdateOrderByIDApp error: ", " | ", fmt.Sprintf("order not found with ID: %v", id))
		return nil, fmt.Errorf("order not found with ID: %v", id)
	}

	var (
		clientUuid       = foundOrder.ClientUUID
		voucherUuid      = foundOrder.VoucherUUID
		status           = foundOrder.Status
		verificationCode = foundOrder.VerificationCode
		createdAt        = foundOrder.CreatedAt
	)

	if len(reqOrder.ClientUUID) > 0 {
		clientUuid = reqOrder.ClientUUID
	}

	if len(reqOrder.VoucherUUID) > 0 {
		voucherUuid = reqOrder.VoucherUUID
	}

	if len(reqOrder.Status) > 0 {
		status = reqOrder.Status
	}

	if len(reqOrder.VerificationCode) > 0 {
		clientUuid = reqOrder.VerificationCode
	}

	if len(reqOrder.CreatedAt) > 0 {
		clientUuid = reqOrder.CreatedAt
	}

	reqOrder = dto.RequestOrder{
		ID:               id,
		ClientUUID:       clientUuid,
		VoucherUUID:      voucherUuid,
		Status:           status,
		VerificationCode: verificationCode,
		CreatedAt:        createdAt,
	}

	if err := app.UpdateOrderByIDUseCase(id, reqOrder); err != nil {
		l.Errorf("UpdateOrderByIDApp error: ", " | ", err)
		return nil, err
	}

	oDB := dto.OrderDB{
		ID:               id,
		ClientUUID:       clientUuid,
		VoucherUUID:      voucherUuid,
		Status:           status,
		VerificationCode: verificationCode,
		CreatedAt:        createdAt,
	}

	order, err := app.UpdateOrderByIDRepository(id, oDB)

	if err != nil {
		l.Errorf("UpdateOrderByIDApp error: ", " | ", err)
		return nil, err
	}

	if order == nil {
		l.Infof("UpdateOrderByIDApp output: ", " | ", nil)
		return nil, errors.New("client is null, is not possible to proceed with update order")
	}

	out := &dto.OutputOrder{
		ID:               order.ID,
		ClientUUID:       order.ClientUUID,
		VoucherUUID:      order.VoucherUUID,
		Status:           order.Status,
		VerificationCode: order.VerificationCode,
		CreatedAt:        order.CreatedAt,
	}

	l.Infof("UpdateOrderByIDApp output: ", " | ", ps.MarshalString(out))
	return out, nil
}

func (app application) GetOrders() ([]dto.OutputOrder, error) {
	l.Infof("GetOrdersApp: ", " | ")
	orderList := make([]dto.OutputOrder, 0)

	if err := app.GetOrdersUseCase(); err != nil {
		l.Errorf("GetOrdersApp error: ", " | ", err)
		return nil, err
	}

	orders, err := app.GetOrdersRepository()

	if err != nil {
		l.Errorf("GetOrdersApp error: ", " | ", err)
		return nil, err

	}

	for i := range orders {

		orderItemList, err := app.GetAllOrderItemByOrderIDRepository(orders[i].ID)

		if err != nil {
			l.Errorf("GetOrdersApp error: ", " | ", err)
			return nil, err
		}

		itemOutList := make([]dto.OutputOrderItem, 0)
		for _, op := range orderItemList {
			itemOutList = append(itemOutList, dto.OutputOrderItem{
				ID:          op.ID,
				OrderID:     op.OrderID,
				ProductUUID: op.ProductUUID,
				Quantity:    op.Quantity,
			})
		}

		order := dto.OutputOrder{
			ID:               orders[i].ID,
			ClientUUID:       orders[i].ClientUUID,
			VoucherUUID:      orders[i].VoucherUUID,
			Items:            itemOutList,
			Status:           orders[i].Status,
			VerificationCode: orders[i].VerificationCode,
			CreatedAt:        orders[i].CreatedAt,
		}

		if strings.ToLower(order.Status) != vo.FinishedStatusKey {
			orderList = append(orderList, order)
		}
	}

	l.Infof("GetOrdersApp output: ", " | ", ps.MarshalString(orderList))
	return orderList, nil
}

func (app application) GetOrderByID(id int64) (*dto.OutputOrder, error) {
	l.Infof("GetOrderByIDApp: ", " | ", id)

	if err := app.GetOrderByIDUseCase(id); err != nil {
		l.Errorf("GetOrderByIDApp error: ", " | ", err)
		return nil, err
	}

	o, err := app.GetOrderByIDRepository(id)

	if err != nil {
		l.Errorf("GetOrderByIDApp error: ", " | ", err)
		return nil, err
	}

	orderItemList, err := app.GetAllOrderItemByOrderIDRepository(id)

	if err != nil {
		l.Errorf("GetOrderByIDApp error: ", " | ", err)
		return nil, err
	}

	itemOutList := make([]dto.OutputOrderItem, 0)
	for _, op := range orderItemList {
		itemOutList = append(itemOutList, dto.OutputOrderItem{
			ID:          op.ID,
			OrderID:     op.OrderID,
			ProductUUID: op.ProductUUID,
			Quantity:    op.Quantity,
		})
	}

	out := &dto.OutputOrder{
		ID:               o.ID,
		ClientUUID:       o.ClientUUID,
		VoucherUUID:      o.VoucherUUID,
		Items:            itemOutList,
		Status:           o.Status,
		VerificationCode: o.VerificationCode,
		CreatedAt:        o.CreatedAt,
	}

	l.Infof("GetOrderByIDApp output: ", " | ", ps.MarshalString(out))
	return out, nil
}

func (app application) SaveOrder(order dto.RequestOrder) (*dto.OutputOrder, error) {
	l.Infof("SaveOrderApp: ", " | ", ps.MarshalString(order))

	if err := app.SaveOrderUseCase(order); err != nil {
		l.Errorf("SaveOrderApp error: ", " | ", err)
		return nil, err
	}

	nowFmt := time.Now().Format("02-01-2006 15:04:05")

	oDb := dto.OrderDB{
		ClientUUID:       order.ClientUUID,
		VoucherUUID:      order.VoucherUUID,
		Items:            order.Items,
		Status:           order.Status,
		VerificationCode: order.VerificationCode,
		CreatedAt:        nowFmt,
	}

	o, err := app.SaveOrderRepository(oDb)

	if err != nil {
		l.Errorf("SaveOrderApp error: ", " | ", err)
		return nil, err

	}

	for _, orderItems := range order.Items {
		opIn := dto.OrderItemDB{
			OrderID:     o.ID,
			ProductUUID: orderItems.ProductUUID,
			Quantity:    orderItems.Quantity,
			TotalPrice:  orderItems.TotalPrice,
			Discount:    orderItems.Discount,
			CreatedAt:   nowFmt,
		}

		opService, err := app.SaveOrderItemRepository(opIn)

		if err != nil {
			l.Errorf("SaveOrderApp error: ", " | ", err)
			return nil, err
		}

		if opService == nil {
			orderProductNullErr := errors.New("is not possible to save order because it's null")
			l.Infof("SaveOrderApp output: ", " | ", orderProductNullErr)
			return nil, orderProductNullErr
		}
	}

	outOrder := &dto.OutputOrder{
		ID:               o.ID,
		ClientUUID:       o.ClientUUID,
		VoucherUUID:      o.VoucherUUID,
		Items:            o.Items,
		Status:           o.Status,
		VerificationCode: o.VerificationCode,
		CreatedAt:        o.CreatedAt,
	}

	l.Infof("SaveOrderApp output: ", " | ", ps.MarshalString(outOrder))
	return outOrder, nil
}
