package rpc

import (
	"context"
	"fase-4-hf-order/internal/core/application"
	"fase-4-hf-order/internal/core/domain/entity"
	"fase-4-hf-order/internal/core/domain/entity/dto"
	cp "fase-4-hf-order/order_proto"
)

type HandlerGRPC interface {
	Handler() *handlerGRPC
}

type handlerGRPC struct {
	app application.Application
	cp.UnimplementedOrderServer
}

func NewHandler(app application.Application) HandlerGRPC {
	return &handlerGRPC{app: app}
}

func (h *handlerGRPC) Handler() *handlerGRPC {
	return h
}

func (h *handlerGRPC) Create(ctx context.Context, cr *cp.CreateRequest) (*cp.CreateResponse, error) {
	items := make([]entity.OrderItems, 0)

	for i := range cr.Items {

		items = append(items, entity.OrderItems{
			ProductUUID: cr.Items[i].ProductUuid,
			Quantity:    cr.Items[i].Quantity,
		})
	}

	input := dto.RequestOrder{
		ClientUUID:  cr.ClientUuid,
		VoucherUUID: cr.ClientUuid,
		Items:       items,
	}

	o, err := h.app.SaveOrder(input)

	if err != nil {
		return nil, err
	}

	if o == nil {
		return nil, nil
	}

	outItems := make([]*cp.Item, 0)

	for i := range o.Items {
		item := cp.Item{
			ProductUuid: o.Items[i].ProductUUID,
			Quantity:    o.Items[i].Quantity,
		}
		outItems = append(outItems, &item)
	}

	out := &cp.CreateResponse{
		Id:               o.ID,
		ClientUuid:       o.ClientUUID,
		VoucherUuid:      o.VoucherUUID,
		Items:            outItems,
		VerificationCode: o.VerificationCode,
		Status:           o.Status,
		CreatedAt:        o.CreatedAt,
	}

	return out, nil
}

func (h *handlerGRPC) GetByID(ctx context.Context, gr *cp.GetByIDRequest) (*cp.GetByIDResponse, error) {
	o, err := h.app.GetOrderByID(gr.Id)

	if err != nil {
		return nil, err
	}

	if o == nil {
		return nil, nil
	}

	outItems := make([]*cp.Item, 0)

	for i := range o.Items {
		item := cp.Item{
			ProductUuid: o.Items[i].ProductUUID,
			Quantity:    o.Items[i].Quantity,
		}
		outItems = append(outItems, &item)
	}

	out := &cp.GetByIDResponse{
		Id:               o.ID,
		ClientUuid:       o.ClientUUID,
		VoucherUuid:      o.VoucherUUID,
		Items:            outItems,
		VerificationCode: o.VerificationCode,
		Status:           o.Status,
		CreatedAt:        o.CreatedAt,
	}

	return out, nil
}

func (h *handlerGRPC) Update(ctx context.Context, ur *cp.UpdateRequest) (*cp.UpdateResponse, error) {
	items := make([]entity.OrderItems, 0)

	for i := range ur.Items {

		items = append(items, entity.OrderItems{
			ProductUUID: ur.Items[i].ProductUuid,
			Quantity:    ur.Items[i].Quantity,
		})
	}

	input := dto.RequestOrder{
		ClientUUID:  ur.ClientUuid,
		VoucherUUID: ur.ClientUuid,
		Items:       items,
	}

	o, err := h.app.UpdateOrderByID(input.ID, input)

	if err != nil {
		return nil, err
	}

	if o == nil {
		return nil, nil
	}

	outItems := make([]*cp.Item, 0)

	for i := range o.Items {
		item := cp.Item{
			ProductUuid: o.Items[i].ProductUUID,
			Quantity:    o.Items[i].Quantity,
		}
		outItems = append(outItems, &item)
	}

	out := &cp.UpdateResponse{
		Id:               o.ID,
		ClientUuid:       o.ClientUUID,
		VoucherUuid:      o.VoucherUUID,
		Items:            outItems,
		VerificationCode: o.VerificationCode,
		Status:           o.Status,
		CreatedAt:        o.CreatedAt,
	}

	return out, nil
}

func (h *handlerGRPC) Get(context.Context, *cp.GetRequest) (*cp.GetResponse, error) {

	o, err := h.app.GetOrders()

	if err != nil {
		return nil, err
	}

	if o == nil {
		return nil, nil
	}

	outOrders := make([]*cp.OrderItem, 0)
	for orderIdx := range o {

		outItems := make([]*cp.Item, 0)

		for i := range o[orderIdx].Items {
			item := cp.Item{
				ProductUuid: o[orderIdx].Items[i].ProductUUID,
				Quantity:    o[orderIdx].Items[i].Quantity,
			}
			outItems = append(outItems, &item)
		}

		oItem := cp.OrderItem{
			Id:               o[orderIdx].ID,
			ClientUuid:       o[orderIdx].ClientUUID,
			VoucherUuid:      o[orderIdx].VoucherUUID,
			Items:            outItems,
			Status:           o[orderIdx].Status,
			VerificationCode: o[orderIdx].VerificationCode,
			CreatedAt:        o[orderIdx].CreatedAt,
		}

		outOrders = append(outOrders, &oItem)

	}

	out := &cp.GetResponse{
		Orders: outOrders,
	}

	return out, nil
}
