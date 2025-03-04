package graph

import (
	"context"

	"go-rest-grpc-graphql-clean-architecture/internal/domain"
)

func (r *queryResolver) ListOrders(ctx context.Context) ([]*domain.Order, error) {
	orders, err := r.listOrdersUseCase.Execute()
	if err != nil {
		return nil, err
	}

	result := make([]*domain.Order, len(orders))
	for i := range orders {
		result[i] = &orders[i]
	}
	return result, nil
}
