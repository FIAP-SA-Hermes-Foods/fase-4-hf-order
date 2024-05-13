package reposql

import (
	"context"
	"encoding/json"
	l "fase-4-hf-order/external/logger"
	ps "fase-4-hf-order/external/strings"
	"fase-4-hf-order/internal/core/db"
	"fase-4-hf-order/internal/core/domain/entity/dto"
	"fase-4-hf-order/internal/core/domain/repository"
)

var (
	queryGetOrderItemByOrderID = `SELECT * from orders_items where orders_id = $1`
	queryGetOrderItems         = `SELECT * from orders_items`
	querySaveOrderItems        = `INSERT INTO orders_items (id, orders_id, product_uuid, quantity, total_price, discount, created_at) VALUES (DEFAULT, $1, $2, $3, $4, $5, now()) RETURNING id, created_at`
)

var _ repository.OrderItemRepository = (*orderItemDB)(nil)

type orderItemDB struct {
	Ctx      context.Context
	Database db.SQLDatabase
}

func NewOrderItemDB(ctx context.Context, sqlDB db.SQLDatabase) *orderItemDB {
	return &orderItemDB{Ctx: ctx, Database: sqlDB}
}

func (o *orderItemDB) GetAllOrderItem() ([]dto.OrderItemDB, error) {
	l.Infof("GetAllOrderItem received input: ", " | ", nil)
	if err := o.Database.Connect(); err != nil {
		l.Errorf("GetAllOrderItem connect error: ", " | ", err)
		return nil, err
	}

	defer o.Database.Close()

	var (
		order     = new(dto.OrderItemDB)
		orderList = make([]dto.OrderItemDB, 0)
	)

	if err := o.Database.Query(queryGetOrderItems); err != nil {
		l.Errorf("GetAllOrderItem error to connect database: ", " | ", err)
		return nil, err
	}

	for o.Database.GetNextRows() {
		var orderItem dto.OrderItemDB

		err := o.Database.Scan(
			&order.ID,
			&order.Quantity,
			&order.TotalPrice,
			&order.Discount,
			&order.OrderID,
			&order.ProductUUID,
			&order.CreatedAt,
		)

		if err != nil {
			l.Errorf("GetAllOrderItem error to scan database: ", " | ", err)
			return nil, err
		}

		orderItem = *order
		orderList = append(orderList, orderItem)
	}

	olStr, err := json.Marshal(orderList)

	if err != nil {
		l.Errorf("GetAllOrderItemByOrderID error to unmarshal: ", " | ", err)
		return nil, err
	}

	l.Infof("GetAllOrderItemByOrderID output: ", " | ", string(olStr))
	return orderList, nil
}

func (o *orderItemDB) GetAllOrderItemByOrderID(id int64) ([]dto.OrderItemDB, error) {
	l.Infof("GetAllOrderItemByOrderID received input: ", " | ", id)
	if err := o.Database.Connect(); err != nil {
		l.Errorf("GetAllOrderItemByOrderID connect error: ", " | ", err)
		return nil, err
	}

	defer o.Database.Close()

	var (
		order     = new(dto.OrderItemDB)
		orderList = make([]dto.OrderItemDB, 0)
	)

	if err := o.Database.Query(queryGetOrderItemByOrderID, id); err != nil {
		l.Errorf("GetAllOrderItemByOrderID error to connect database: ", " | ", err)
		return nil, err
	}

	for o.Database.GetNextRows() {
		var orderItem dto.OrderItemDB

		err := o.Database.Scan(
			&order.ID,
			&order.Quantity,
			&order.TotalPrice,
			&order.Discount,
			&order.OrderID,
			&order.ProductUUID,
			&order.CreatedAt,
		)

		if err != nil {
			l.Errorf("GetAllOrderItemByOrderID error to scan database: ", " | ", err)
			return nil, err
		}

		orderItem = *order
		orderList = append(orderList, orderItem)
	}

	olStr, err := json.Marshal(orderList)

	if err != nil {
		l.Errorf("GetAllOrderItemByOrderID error to unmarshal: ", " | ", err)
		return nil, err
	}

	l.Infof("GetAllOrderItemByOrderID output: ", " | ", string(olStr))
	return orderList, nil
}

func (o *orderItemDB) SaveOrderItem(order dto.OrderItemDB) (*dto.OrderItemDB, error) {
	l.Infof("SaveOrderItem received input: ", " | ", order)
	if err := o.Database.Connect(); err != nil {
		l.Errorf("SaveOrderItem connect error: ", " | ", err)
		return nil, err
	}

	defer o.Database.Close()

	if err := o.Database.PrepareStmt(querySaveOrderItems); err != nil {
		l.Errorf("SaveOrderItem error to prepare statement: ", " | ", err)
		return nil, err
	}

	defer o.Database.CloseStmt()

	var outOrder = &dto.OrderItemDB{
		Quantity:    order.ID,
		TotalPrice:  order.TotalPrice,
		Discount:    order.Discount,
		OrderID:     order.OrderID,
		ProductUUID: order.ProductUUID,
	}

	o.Database.QueryRow(order.OrderID, order.ProductUUID, order.Quantity, order.TotalPrice, order.Discount)

	if err := o.Database.ScanStmt(&outOrder.ID, &outOrder.CreatedAt); err != nil {
		l.Errorf("SaveOrderItem error to scan statement: ", " | ", err)
		return nil, err
	}

	l.Infof("SaveOrderItem output: ", " | ", ps.MarshalString(outOrder))
	return outOrder, nil
}
