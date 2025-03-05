package usecase

import "go-rest-grpc-graphql-clean-architecture/internal/domain"

// ListOrdersUseCase define a interface para listar pedidos
type ListOrdersUseCase interface {
	Execute() ([]domain.Order, error)
}

// CreateOrderUseCase define a interface para criar pedidos
type CreateOrderUseCase interface {
	Execute(order *domain.Order) error
}
