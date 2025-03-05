package rest

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"

	"go-rest-grpc-graphql-clean-architecture/internal/domain"
	"go-rest-grpc-graphql-clean-architecture/internal/usecase"
)

// OrderHandler gerencia as requisições HTTP relacionadas a pedidos
type OrderHandler struct {
	listOrdersUseCase  usecase.ListOrdersUseCase
	createOrderUseCase usecase.CreateOrderUseCase
}

// NewOrderHandler cria uma nova instância do handler
func NewOrderHandler(listOrdersUseCase usecase.ListOrdersUseCase, createOrderUseCase usecase.CreateOrderUseCase) *OrderHandler {
	return &OrderHandler{
		listOrdersUseCase:  listOrdersUseCase,
		createOrderUseCase: createOrderUseCase,
	}
}

// ListOrders retorna a lista de todos os pedidos
func (h *OrderHandler) ListOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.listOrdersUseCase.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

// CreateOrder cria um novo pedido
func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order domain.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order.ID = uuid.New().String()
	if err := h.createOrderUseCase.Execute(&order); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}
