package repository

import (
	"go-rest-grpc-graphql-clean-architecture/internal/domain"

	"gorm.io/gorm"
)

// OrderRepository implementa a interface de repositório para pedidos
type OrderRepository struct {
	db *gorm.DB
}

// NewOrderRepository cria uma nova instância do repositório
func NewOrderRepository(db *gorm.DB) domain.OrderRepository {
	return &OrderRepository{db: db}
}

// Create salva um novo pedido no banco de dados
func (r *OrderRepository) Create(order *domain.Order) error {
	return r.db.Create(order).Error
}

// List retorna todos os pedidos do banco de dados
func (r *OrderRepository) List() ([]domain.Order, error) {
	var orders []domain.Order
	err := r.db.Find(&orders).Error
	return orders, err
}
