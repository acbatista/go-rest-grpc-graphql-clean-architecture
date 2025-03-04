package graph

import (
	"context"
	"time"

	"go-rest-grpc-graphql-clean-architecture/internal/domain"
)

func (r *orderResolver) CreatedAt(ctx context.Context, obj *domain.Order) (string, error) {
	return obj.CreatedAt.Format(time.RFC3339), nil
}

func (r *orderResolver) UpdatedAt(ctx context.Context, obj *domain.Order) (string, error) {
	return obj.UpdatedAt.Format(time.RFC3339), nil
}
