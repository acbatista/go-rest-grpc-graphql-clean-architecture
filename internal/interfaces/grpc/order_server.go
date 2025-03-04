package grpc

import (
	"context"

	"go-rest-grpc-graphql-clean-architecture/internal/usecase"
	pb "go-rest-grpc-graphql-clean-architecture/proto"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type OrderServer struct {
	pb.UnimplementedOrderServiceServer
	listOrdersUseCase *usecase.ListOrdersUseCase
}

func NewOrderServer(listOrdersUseCase *usecase.ListOrdersUseCase) *OrderServer {
	return &OrderServer{
		listOrdersUseCase: listOrdersUseCase,
	}
}

func (s *OrderServer) ListOrders(ctx context.Context, req *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	orders, err := s.listOrdersUseCase.Execute()
	if err != nil {
		return nil, err
	}

	var pbOrders []*pb.Order
	for _, order := range orders {
		pbOrders = append(pbOrders, &pb.Order{
			Id:           order.ID,
			CustomerName: order.CustomerName,
			Total:        order.Total,
			Status:       order.Status,
			CreatedAt:    timestamppb.New(order.CreatedAt),
			UpdatedAt:    timestamppb.New(order.UpdatedAt),
		})
	}

	return &pb.ListOrdersResponse{
		Orders: pbOrders,
	}, nil
}
