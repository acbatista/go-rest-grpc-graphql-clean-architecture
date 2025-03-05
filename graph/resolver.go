package graph

import (
	"go-rest-grpc-graphql-clean-architecture/internal/interfaces/graphql"
	"go-rest-grpc-graphql-clean-architecture/internal/usecase"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

func NewResolver(listOrdersUseCase usecase.ListOrdersUseCase, createOrderUseCase usecase.CreateOrderUseCase) *graphql.Resolver {
	return graphql.NewResolver(listOrdersUseCase, createOrderUseCase)
}
