package domain

import (
	"errors"
	"time"
)

// Order representa um pedido no sistema
type Order struct {
	ID           string    `json:"id" gorm:"primaryKey"`
	CustomerName string    `json:"customer_name"`
	Total        float64   `json:"total"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// OrderRepository define a interface para operações com pedidos
type OrderRepository interface {
	List() ([]Order, error)
	Create(order *Order) error
}

// Validate valida os campos do pedido
func (o *Order) Validate() error {
	if o.CustomerName == "" {
		return errors.New("nome do cliente é obrigatório")
	}
	if o.Total <= 0 {
		return errors.New("total deve ser maior que zero")
	}
	if o.Status == "" {
		return errors.New("status é obrigatório")
	}
	return nil
}
