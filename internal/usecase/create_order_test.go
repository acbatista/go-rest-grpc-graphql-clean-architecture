package usecase

import (
	"testing"

	"go-rest-grpc-graphql-clean-architecture/internal/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateOrderUseCase_Execute(t *testing.T) {
	mockRepo := new(MockOrderRepository)
	useCase := NewCreateOrderUseCase(mockRepo)

	order := &domain.Order{
		ID:           "1",
		CustomerName: "John Doe",
		Total:        100.0,
		Status:       "pending",
	}

	mockRepo.On("Create", mock.AnythingOfType("*domain.Order")).Return(nil)

	err := useCase.Execute(order)

	assert.NoError(t, err)
	assert.NotZero(t, order.CreatedAt)
	assert.NotZero(t, order.UpdatedAt)
	mockRepo.AssertExpectations(t)
}

func TestCreateOrderUseCase_Execute_Error(t *testing.T) {
	mockRepo := new(MockOrderRepository)
	useCase := NewCreateOrderUseCase(mockRepo)

	order := &domain.Order{
		ID:           "1",
		CustomerName: "John Doe",
		Total:        100.0,
		Status:       "pending",
	}

	mockRepo.On("Create", mock.AnythingOfType("*domain.Order")).Return(assert.AnError)

	err := useCase.Execute(order)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCreateOrderUseCase_Execute_ValidationError(t *testing.T) {
	mockRepo := new(MockOrderRepository)
	useCase := NewCreateOrderUseCase(mockRepo)

	order := &domain.Order{
		ID:           "1",
		CustomerName: "", // Invalid: empty customer name
		Total:        100.0,
		Status:       "pending",
	}

	err := useCase.Execute(order)

	assert.Error(t, err)
	mockRepo.AssertNotCalled(t, "Create")
}
