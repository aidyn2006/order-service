package domain

import "context"

type OrderItem struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	OrderID   uint    `gorm:"not null" json:"order_id"`
	ProductID uint    `gorm:"not null" json:"product_id"`
	Quantity  int     `gorm:"not null" json:"quantity"`
	Price     float64 `gorm:"not null" json:"price"`
}

type OrderItemRepository interface {
	CreateOrderItem(ctx context.Context, item *OrderItem) error
	GetOrderItemsByOrderID(ctx context.Context, orderID uint) ([]OrderItem, error)
	UpdateOrderItem(ctx context.Context, item *OrderItem) error
	DeleteOrderItem(ctx context.Context, id uint) error
}

type OrderItemUseCase interface {
	CreateOrderItem(ctx context.Context, item *OrderItem) error
	GetOrderItemsByOrderID(ctx context.Context, orderID uint) ([]OrderItem, error)
	UpdateOrderItem(ctx context.Context, item *OrderItem) error
	DeleteOrderItem(ctx context.Context, id uint) error
}
