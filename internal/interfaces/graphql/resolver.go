package graphql

import (
	"context"

	"go-rest-grpc-graphql-clean-architecture/graph/generated"
	"go-rest-grpc-graphql-clean-architecture/internal/domain"
	"go-rest-grpc-graphql-clean-architecture/internal/usecase"
)

type Resolver struct {
	listOrdersUseCase  usecase.ListOrdersUseCase
	createOrderUseCase usecase.CreateOrderUseCase
}

func NewResolver(listOrdersUseCase usecase.ListOrdersUseCase, createOrderUseCase usecase.CreateOrderUseCase) *Resolver {
	return &Resolver{
		listOrdersUseCase:  listOrdersUseCase,
		createOrderUseCase: createOrderUseCase,
	}
}

func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) Order() generated.OrderResolver {
	return &orderResolver{r}
}

type queryResolver struct {
	*Resolver
}

func (r *queryResolver) ListOrders(ctx context.Context) ([]*domain.Order, error) {
	orders, err := r.listOrdersUseCase.Execute()
	if err != nil {
		return nil, err
	}

	var result []*domain.Order
	for i := range orders {
		result = append(result, &orders[i])
	}
	return result, nil
}

type orderResolver struct {
	*Resolver
}

func (r *orderResolver) CreatedAt(ctx context.Context, obj *domain.Order) (string, error) {
	return obj.CreatedAt.Format("2006-01-02T15:04:05Z07:00"), nil
}

func (r *orderResolver) UpdatedAt(ctx context.Context, obj *domain.Order) (string, error) {
	return obj.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"), nil
}
