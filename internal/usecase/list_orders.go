package usecase

import (
	"go-rest-grpc-graphql-clean-architecture/internal/domain"
)

// ListOrdersUseCaseImpl implementa o caso de uso para listar pedidos
type ListOrdersUseCaseImpl struct {
	orderRepository domain.OrderRepository
}

// NewListOrdersUseCase cria uma nova instância do caso de uso
func NewListOrdersUseCase(orderRepository domain.OrderRepository) ListOrdersUseCase {
	return &ListOrdersUseCaseImpl{
		orderRepository: orderRepository,
	}
}

// Execute executa o caso de uso de listagem de pedidos
func (u *ListOrdersUseCaseImpl) Execute() ([]domain.Order, error) {
	return u.orderRepository.List()
}
