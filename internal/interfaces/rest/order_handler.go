package rest

import (
	"encoding/json"
	"net/http"

	"go-rest-grpc-graphql-clean-architecture/internal/domain"
	"go-rest-grpc-graphql-clean-architecture/internal/usecase"

	"github.com/google/uuid"
)

type OrderHandler struct {
	listOrdersUseCase  *usecase.ListOrdersUseCase
	createOrderUseCase *usecase.CreateOrderUseCase
}

type CreateOrderRequest struct {
	CustomerName string  `json:"customer_name"`
	Total        float64 `json:"total"`
	Status       string  `json:"status"`
}

func NewOrderHandler(listOrdersUseCase *usecase.ListOrdersUseCase, createOrderUseCase *usecase.CreateOrderUseCase) *OrderHandler {
	return &OrderHandler{
		listOrdersUseCase:  listOrdersUseCase,
		createOrderUseCase: createOrderUseCase,
	}
}

func (h *OrderHandler) ListOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.listOrdersUseCase.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	order := &domain.Order{
		ID:           uuid.New().String(),
		CustomerName: req.CustomerName,
		Total:        req.Total,
		Status:       req.Status,
	}

	if err := h.createOrderUseCase.Execute(order); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}
