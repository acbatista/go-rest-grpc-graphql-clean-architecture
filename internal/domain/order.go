package domain

import (
	"time"
)

type Order struct {
	ID           string    `json:"id"`
	CustomerName string    `json:"customer_name"`
	Total        float64   `json:"total"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type OrderRepository interface {
	List() ([]Order, error)
	Create(order *Order) error
}
