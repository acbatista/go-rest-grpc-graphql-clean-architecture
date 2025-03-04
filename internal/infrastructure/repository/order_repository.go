package repository

import (
	"go-rest-grpc-graphql-clean-architecture/internal/domain"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) List() ([]domain.Order, error) {
	var orders []domain.Order
	result := r.db.Find(&orders)
	return orders, result.Error
}

func (r *OrderRepository) Create(order *domain.Order) error {
	return r.db.Create(order).Error
}
