package domain

import (
	"gorm.io/gorm"
	"time"
)

type OrderStatus string

const (
	StatusPending   OrderStatus = "pending"
	StatusCompleted OrderStatus = "completed"
	StatusCanceled  OrderStatus = "canceled"
)

type Order struct {
	gorm.Model
	UserID    uint        `json:"user_id"`
	Total     float64     `json:"total"`
	Status    OrderStatus `json:"status"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Items     []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
}
