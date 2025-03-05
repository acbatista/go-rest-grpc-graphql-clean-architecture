package grpcserver

import (
	"context"

	"go-rest-grpc-graphql-clean-architecture/internal/usecase"
	pb "go-rest-grpc-graphql-clean-architecture/proto"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// OrderServer implementa o serviço gRPC de pedidos
type OrderServer struct {
	pb.UnimplementedOrderServiceServer
	listOrdersUseCase usecase.ListOrdersUseCase
}

// NewOrderServer cria uma nova instância do servidor gRPC
func NewOrderServer(listOrdersUseCase usecase.ListOrdersUseCase) *OrderServer {
	return &OrderServer{
		listOrdersUseCase: listOrdersUseCase,
	}
}

// ListOrders retorna a lista de todos os pedidos
func (s *OrderServer) ListOrders(ctx context.Context, req *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	orders, err := s.listOrdersUseCase.Execute()
	if err != nil {
		return nil, err
	}

	var responseOrders []*pb.Order
	for _, order := range orders {
		responseOrders = append(responseOrders, &pb.Order{
			Id:           order.ID,
			CustomerName: order.CustomerName,
			Total:        order.Total,
			Status:       order.Status,
			CreatedAt:    timestamppb.New(order.CreatedAt),
			UpdatedAt:    timestamppb.New(order.UpdatedAt),
		})
	}

	return &pb.ListOrdersResponse{
		Orders: responseOrders,
	}, nil
}
