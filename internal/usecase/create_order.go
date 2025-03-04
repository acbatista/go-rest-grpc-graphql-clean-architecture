package usecase

import (
	"time"

	"go-rest-grpc-graphql-clean-architecture/internal/domain"
)

type CreateOrderUseCase struct {
	repo domain.OrderRepository
}

func NewCreateOrderUseCase(repo domain.OrderRepository) *CreateOrderUseCase {
	return &CreateOrderUseCase{repo: repo}
}

func (uc *CreateOrderUseCase) Execute(order *domain.Order) error {
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()
	return uc.repo.Create(order)
}
