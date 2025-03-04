package graphql

import (
	"context"

	"go-rest-grpc-graphql-clean-architecture/internal/domain"
	"go-rest-grpc-graphql-clean-architecture/internal/usecase"
)

type Resolver struct {
	listOrdersUseCase *usecase.ListOrdersUseCase
}

func NewResolver(listOrdersUseCase *usecase.ListOrdersUseCase) *Resolver {
	return &Resolver{
		listOrdersUseCase: listOrdersUseCase,
	}
}

type queryResolver struct {
	*Resolver
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type QueryResolver interface {
	ListOrders(context.Context) ([]domain.Order, error)
}

func (r *queryResolver) ListOrders(ctx context.Context) ([]domain.Order, error) {
	return r.listOrdersUseCase.Execute()
}
