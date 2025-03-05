package rest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-rest-grpc-graphql-clean-architecture/internal/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockListOrdersUseCase struct {
	mock.Mock
}

func (m *MockListOrdersUseCase) Execute() ([]domain.Order, error) {
	args := m.Called()
	return args.Get(0).([]domain.Order), args.Error(1)
}

type MockCreateOrderUseCase struct {
	mock.Mock
}

func (m *MockCreateOrderUseCase) Execute(order *domain.Order) error {
	args := m.Called(order)
	return args.Error(0)
}

func TestOrderHandler_ListOrders(t *testing.T) {
	mockListUseCase := new(MockListOrdersUseCase)
	mockCreateUseCase := new(MockCreateOrderUseCase)
	handler := NewOrderHandler(mockListUseCase, mockCreateUseCase)

	expectedOrders := []domain.Order{
		{
			ID:           "1",
			CustomerName: "John Doe",
			Total:        100.0,
			Status:       "pending",
		},
	}

	mockListUseCase.On("Execute").Return(expectedOrders, nil)

	req := httptest.NewRequest("GET", "/order", nil)
	w := httptest.NewRecorder()

	handler.ListOrders(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var response []domain.Order
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, expectedOrders, response)
}

func TestOrderHandler_CreateOrder(t *testing.T) {
	mockListUseCase := new(MockListOrdersUseCase)
	mockCreateUseCase := new(MockCreateOrderUseCase)
	handler := NewOrderHandler(mockListUseCase, mockCreateUseCase)

	order := &domain.Order{
		CustomerName: "John Doe",
		Total:        100.0,
		Status:       "pending",
	}

	body, _ := json.Marshal(order)
	req := httptest.NewRequest("POST", "/order", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	mockCreateUseCase.On("Execute", mock.AnythingOfType("*domain.Order")).Return(nil)

	handler.CreateOrder(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var response domain.Order
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.NotEmpty(t, response.ID)
	assert.Equal(t, order.CustomerName, response.CustomerName)
	assert.Equal(t, order.Total, response.Total)
	assert.Equal(t, order.Status, response.Status)
}

func TestOrderHandler_CreateOrder_InvalidRequest(t *testing.T) {
	mockListUseCase := new(MockListOrdersUseCase)
	mockCreateUseCase := new(MockCreateOrderUseCase)
	handler := NewOrderHandler(mockListUseCase, mockCreateUseCase)

	invalidBody := []byte(`{"invalid": "json"`)
	req := httptest.NewRequest("POST", "/order", bytes.NewBuffer(invalidBody))
	w := httptest.NewRecorder()

	handler.CreateOrder(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockCreateUseCase.AssertNotCalled(t, "Execute")
}
