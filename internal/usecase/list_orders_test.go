package usecase

import (
	"testing"
	"time"

	"go-rest-grpc-graphql-clean-architecture/internal/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockOrderRepository struct {
	mock.Mock
}

func (m *MockOrderRepository) List() ([]domain.Order, error) {
	args := m.Called()
	return args.Get(0).([]domain.Order), args.Error(1)
}

func (m *MockOrderRepository) Create(order *domain.Order) error {
	args := m.Called(order)
	return args.Error(0)
}

func TestListOrdersUseCase_Execute(t *testing.T) {
	mockRepo := new(MockOrderRepository)
	useCase := NewListOrdersUseCase(mockRepo)

	expectedOrders := []domain.Order{
		{
			ID:           "1",
			CustomerName: "John Doe",
			Total:        100.0,
			Status:       "pending",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			ID:           "2",
			CustomerName: "Jane Doe",
			Total:        200.0,
			Status:       "completed",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
	}

	mockRepo.On("List").Return(expectedOrders, nil)

	orders, err := useCase.Execute()

	assert.NoError(t, err)
	assert.Equal(t, expectedOrders, orders)
	mockRepo.AssertExpectations(t)
}

func TestListOrdersUseCase_Execute_Error(t *testing.T) {
	mockRepo := new(MockOrderRepository)
	useCase := NewListOrdersUseCase(mockRepo)

	mockRepo.On("List").Return([]domain.Order{}, assert.AnError)

	orders, err := useCase.Execute()

	assert.Error(t, err)
	assert.Empty(t, orders)
	mockRepo.AssertExpectations(t)
}
