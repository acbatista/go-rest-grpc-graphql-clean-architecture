package usecase

import (
	"go-rest-grpc-graphql-clean-architecture/internal/domain"
	"time"
)

// CreateOrderUseCaseImpl implementa o caso de uso para criar um pedido
type CreateOrderUseCaseImpl struct {
	orderRepository domain.OrderRepository
}

// NewCreateOrderUseCase cria uma nova instância do caso de uso
func NewCreateOrderUseCase(orderRepository domain.OrderRepository) CreateOrderUseCase {
	return &CreateOrderUseCaseImpl{
		orderRepository: orderRepository,
	}
}

// Execute executa o caso de uso de criação de pedido
func (u *CreateOrderUseCaseImpl) Execute(order *domain.Order) error {
	if err := order.Validate(); err != nil {
		return err
	}

	now := time.Now()
	order.CreatedAt = now
	order.UpdatedAt = now

	return u.orderRepository.Create(order)
}
