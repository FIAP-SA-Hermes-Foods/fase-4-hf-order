package http

import (
	"bytes"
	"encoding/json"
	ps "fase-4-hf-order/external/strings"
	"fase-4-hf-order/internal/core/application"
	"fase-4-hf-order/internal/core/domain/entity/dto"
	"fmt"
	"net/http"
	"strconv"
)

type OrderHandler interface {
	Handler(rw http.ResponseWriter, req *http.Request)
	HealthCheck(rw http.ResponseWriter, req *http.Request)
}

type orderHandler struct {
	app application.Application
}

func NewHandler(app application.Application) OrderHandler {
	return orderHandler{app: app}
}

func (h orderHandler) Handler(rw http.ResponseWriter, req *http.Request) {
	var routesOrders = map[string]http.HandlerFunc{
		"get hermes_foods/order":        h.getOrders,
		"get hermes_foods/order/{id}":   h.getOrderByID,
		"post hermes_foods/order":       h.saveOrder,
		"patch hermes_foods/order/{id}": h.updateOrderByID,
	}

	handler, err := router(req.Method, req.URL.Path, routesOrders)

	if err == nil {
		handler(rw, req)
		return
	}

	rw.WriteHeader(http.StatusNotFound)
	rw.Write([]byte(`{"error": "route ` + req.Method + " " + req.URL.Path + ` not found"} `))
}

func (h orderHandler) HealthCheck(rw http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		rw.Write([]byte(`{"error": "method not allowed"}`))
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(`{"status": "OK"}`))
}

func (h orderHandler) saveOrder(rw http.ResponseWriter, req *http.Request) {
	var buff bytes.Buffer
	var reqOrder dto.RequestOrder

	if _, err := buff.ReadFrom(req.Body); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to read data body: %v"} `, err)
		return
	}

	if err := json.Unmarshal(buff.Bytes(), &reqOrder); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to Unmarshal: %v"} `, err)
		return
	}

	o, err := h.app.SaveOrder(reqOrder)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to save order: %v"} `, err)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte(ps.MarshalString(o)))
}

func (h orderHandler) getOrderByID(rw http.ResponseWriter, req *http.Request) {
	id := getID("order", req.URL.Path)

	idconv, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to get order by ID: %v"} `, err)
		return
	}

	o, err := h.app.GetOrderByID(idconv)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to get order by ID: %v"} `, err)
		return
	}

	if o == nil {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte(`{"error": "order not found"}`))
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(ps.MarshalString(o)))
}

func (h orderHandler) getOrders(rw http.ResponseWriter, req *http.Request) {
	oList, err := h.app.GetOrders()

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to get order by ID: %v"} `, err)
		return
	}

	if oList == nil {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte(`{"error": "order not found"}`))
		return
	}

	b, err := json.Marshal(oList)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to get order by ID: %v"} `, err)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write(b)
}

func (h orderHandler) updateOrderByID(rw http.ResponseWriter, req *http.Request) {
	id := getID("order", req.URL.Path)

	idconv, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to get order by ID: %v"} `, err)
		return
	}

	var buff bytes.Buffer

	var reqOrder dto.RequestOrder

	if _, err := buff.ReadFrom(req.Body); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to read data body: %v"} `, err)
		return
	}

	if err := json.Unmarshal(buff.Bytes(), &reqOrder); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to Unmarshal: %v"} `, err)
		return
	}

	o, err := h.app.UpdateOrderByID(idconv, reqOrder)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, `{"error": "error to get order by ID: %v"} `, err)
		return
	}

	if o == nil {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte(`{"error": "order not found"}`))
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(ps.MarshalString(o)))
}
