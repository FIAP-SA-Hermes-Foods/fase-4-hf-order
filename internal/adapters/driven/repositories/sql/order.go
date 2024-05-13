package reposql

import (
	"context"
	"encoding/json"
	l "fase-4-hf-order/external/logger"
	ps "fase-4-hf-order/external/strings"
	"fase-4-hf-order/internal/core/db"
	"fase-4-hf-order/internal/core/domain/entity/dto"
	"fase-4-hf-order/internal/core/domain/repository"
	"reflect"
)

var (
	queryGetOrders    = `SELECT * FROM orders ORDER BY created_at ASC`
	queryGetOrderByID = `SELECT * FROM orders WHERE id = $1`
	querySaveOrder    = `INSERT INTO orders (id, status, verification_code, created_at, client_uuid, voucher_uuid) VALUES (DEFAULT, $1, $2, now(), $3, $4) RETURNING id, created_at`
	queryUpdateOrder  = `UPDATE orders SET status = $1, client_uuid = $2, voucher_uuid = $3 WHERE id = $4 RETURNING id, created_at`
)

var _ repository.OrderRepository = (*orderDB)(nil)

type orderDB struct {
	Ctx      context.Context
	Database db.SQLDatabase
}

func NewOrderDB(ctx context.Context, sqlDB db.SQLDatabase) *orderDB {
	return &orderDB{Ctx: ctx, Database: sqlDB}
}

func (o orderDB) SaveOrder(order dto.OrderDB) (*dto.OrderDB, error) {
	l.Infof("SaveOrder received input: ", " | ", order)
	if err := o.Database.Connect(); err != nil {
		l.Errorf("SaveOrder connect error: ", " | ", err)
		return nil, err
	}

	defer o.Database.Close()

	if err := o.Database.PrepareStmt(querySaveOrder); err != nil {
		l.Errorf("SaveOrder prepare error: ", " | ", err)
		return nil, err
	}

	defer o.Database.CloseStmt()

	var outOrder = &dto.OrderDB{
		ClientUUID:       order.ClientUUID,
		VoucherUUID:      order.VoucherUUID,
		Status:           order.Status,
		VerificationCode: order.VerificationCode,
	}

	o.Database.QueryRow(order.Status, order.VerificationCode, order.ClientUUID, order.VoucherUUID)

	if err := o.Database.ScanStmt(&outOrder.ID, &outOrder.CreatedAt); err != nil {
		l.Errorf("SaveOrder scan error: ", " | ", err)
		return nil, err
	}

	l.Infof("SaveOrder output: ", " | ", ps.MarshalString(outOrder))
	return outOrder, nil
}

func (o orderDB) UpdateOrderByID(id int64, order dto.OrderDB) (*dto.OrderDB, error) {
	l.Infof("UpdateOrderByID received input: ", " | ", order)
	if err := o.Database.Connect(); err != nil {
		l.Errorf("UpdateOrderByID connect error: ", " | ", err)
		return nil, err
	}

	defer o.Database.Close()

	if err := o.Database.PrepareStmt(queryUpdateOrder); err != nil {
		l.Errorf("UpdateOrderByID prepare error: ", " | ", err)
		return nil, err
	}

	defer o.Database.CloseStmt()

	var outOrder = &dto.OrderDB{
		ClientUUID:       order.ClientUUID,
		VoucherUUID:      order.VoucherUUID,
		Status:           order.Status,
		VerificationCode: order.VerificationCode,
	}

	o.Database.QueryRow(order.Status, order.ClientUUID, order.VoucherUUID, id)

	if err := o.Database.ScanStmt(&outOrder.ID, &outOrder.CreatedAt); err != nil {
		l.Errorf("UpdateOrderByID scan error: ", " | ", err)
		return nil, err
	}

	l.Infof("UpdateOrderByID output: ", " | ", ps.MarshalString(outOrder))
	return outOrder, nil
}

func (o orderDB) GetOrderByID(id int64) (*dto.OrderDB, error) {
	l.Infof("GetOrderByID received input: ", " | ", id)
	if err := o.Database.Connect(); err != nil {
		l.Errorf("GetOrderByID connect error: ", " | ", err)
		return nil, err
	}

	defer o.Database.Close()

	var outOrder = new(dto.OrderDB)

	if err := o.Database.Query(queryGetOrderByID, id); err != nil {
		l.Errorf("GetOrderByID query error: ", " | ", err)
		return nil, err
	}

	for o.Database.GetNextRows() {
		err := o.Database.Scan(
			&outOrder.ID,
			&outOrder.Status,
			&outOrder.VerificationCode,
			&outOrder.CreatedAt,
			&outOrder.ClientUUID,
			&outOrder.VoucherUUID,
		)
		if err != nil {
			l.Errorf("GetOrderByID scan error: ", " | ", err)
			return nil, err
		}
	}

	if reflect.ValueOf(outOrder).IsNil() || reflect.ValueOf(outOrder).IsZero() {
		return nil, nil
	}

	l.Infof("GetOrderByID output: ", " | ", ps.MarshalString(outOrder))
	return outOrder, nil
}

func (o orderDB) GetOrders() ([]dto.OrderDB, error) {
	l.Infof("GetOrders received input: ", " | ", nil)
	if err := o.Database.Connect(); err != nil {
		l.Errorf("GetOrders connect error: ", " | ", err)
		return nil, err
	}

	defer o.Database.Close()

	var (
		order     = new(dto.OrderDB)
		orderList = make([]dto.OrderDB, 0)
	)

	if err := o.Database.Query(queryGetOrders); err != nil {
		l.Errorf("GetOrders query error: ", " | ", err)
		return nil, err
	}

	for o.Database.GetNextRows() {
		var orderItem dto.OrderDB

		err := o.Database.Scan(
			&order.ID,
			&order.Status,
			&order.VerificationCode,
			&order.CreatedAt,
			&order.ClientUUID,
			&order.VoucherUUID,
		)

		if err != nil {
			l.Errorf("GetOrders scan error: ", " | ", err)
			return nil, err
		}

		orderItem = *order
		orderList = append(orderList, orderItem)
	}

	olStr, err := json.Marshal(orderList)

	if err != nil {
		l.Errorf("GetOrders error to unmarshal: ", " | ", err)
		return nil, err
	}

	l.Infof("GetOrders output: ", " | ", string(olStr))
	return orderList, nil
}
