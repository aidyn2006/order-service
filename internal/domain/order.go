package domain

import (
	"context"
	"errors"
	"time"
)

type OrderStatus string

const (
	StatusPending  OrderStatus = "pending"
	StatusPaid     OrderStatus = "paid"
	StatusShipped  OrderStatus = "shipped"
	StatusCanceled OrderStatus = "canceled"
)

var (
	ErrOrderNotFound = errors.New("order not found")
)

type Order struct {
	ID        uint        `gorm:"primaryKey" json:"id"`
	UserID    uint        `gorm:"not null" json:"user_id"`
	Status    OrderStatus `gorm:"type:varchar(20);not null" json:"status"`
	Total     float64     `gorm:"not null" json:"total"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Items     []OrderItem `gorm:"foreignKey:OrderID" json:"items"`
}

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *Order) error
	GetOrderByID(ctx context.Context, id uint) (*Order, error)
	UpdateOrder(ctx context.Context, order *Order) error
	ListOrders(ctx context.Context, userID uint) ([]Order, error)
	CancelOrder(ctx context.Context, id uint) error
}

type OrderUseCase interface {
	CreateOrder(ctx context.Context, order *Order) error
	GetOrderByID(ctx context.Context, id uint) (*Order, error)
	UpdateOrderStatus(ctx context.Context, order *Order) error
	ListOrders(ctx context.Context, userID uint) ([]Order, error)
	CancelOrder(ctx context.Context, id uint) error
	CalculateOrderTotal(ctx context.Context, order *Order) (float64, error)
}
