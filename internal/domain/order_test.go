package domain

import (
	"testing"
	"time"
)

func TestOrder_Validate(t *testing.T) {
	tests := []struct {
		name    string
		order   *Order
		wantErr bool
	}{
		{
			name: "valid order",
			order: &Order{
				ID:           "123",
				CustomerName: "John Doe",
				Total:        100.0,
				Status:       "pending",
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			},
			wantErr: false,
		},
		{
			name: "invalid order - empty customer name",
			order: &Order{
				ID:           "123",
				CustomerName: "",
				Total:        100.0,
				Status:       "pending",
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			},
			wantErr: true,
		},
		{
			name: "invalid order - negative total",
			order: &Order{
				ID:           "123",
				CustomerName: "John Doe",
				Total:        -100.0,
				Status:       "pending",
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			},
			wantErr: true,
		},
		{
			name: "invalid order - empty status",
			order: &Order{
				ID:           "123",
				CustomerName: "John Doe",
				Total:        100.0,
				Status:       "",
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.order.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Order.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
