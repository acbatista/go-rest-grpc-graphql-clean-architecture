package graph

import "go-rest-grpc-graphql-clean-architecture/internal/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	listOrdersUseCase  *usecase.ListOrdersUseCase
	createOrderUseCase *usecase.CreateOrderUseCase
}

func NewResolver(listOrdersUseCase *usecase.ListOrdersUseCase, createOrderUseCase *usecase.CreateOrderUseCase) *Resolver {
	return &Resolver{
		listOrdersUseCase:  listOrdersUseCase,
		createOrderUseCase: createOrderUseCase,
	}
}
