package usecase

import (
	"go-rest-grpc-graphql-clean-architecture/internal/domain"
)

type ListOrdersUseCase struct {
	repo domain.OrderRepository
}

func NewListOrdersUseCase(repo domain.OrderRepository) *ListOrdersUseCase {
	return &ListOrdersUseCase{repo: repo}
}

func (uc *ListOrdersUseCase) Execute() ([]domain.Order, error) {
	return uc.repo.List()
}
