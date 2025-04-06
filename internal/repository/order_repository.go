package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"order-service/internal/domain"
)

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) domain.OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) CreateOrder(ctx context.Context, order *domain.Order) error {
	return r.db.WithContext(ctx).Create(order).Error
}

func (r *orderRepository) GetOrderByID(ctx context.Context, id uint) (*domain.Order, error) {
	var order domain.Order
	err := r.db.WithContext(ctx).Preload("Items").First(&order, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domain.ErrOrderNotFound
	}
	return &order, err
}

func (r *orderRepository) UpdateOrder(ctx context.Context, order *domain.Order) error {
	result := r.db.WithContext(ctx).Model(&domain.Order{}).Where("id = ?", order.ID).Updates(order)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domain.ErrOrderNotFound
	}
	return nil
}

func (r *orderRepository) ListOrders(ctx context.Context, userID uint) ([]domain.Order, error) {
	var orders []domain.Order
	query := r.db.WithContext(ctx).Preload("Items")

	if userID != 0 {
		query = query.Where("user_id = ?", userID)
	}

	err := query.Find(&orders).Error
	return orders, err
}

func (r *orderRepository) CancelOrder(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).Model(&domain.Order{}).
		Where("id = ?", id).
		Update("status", domain.StatusCanceled)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domain.ErrOrderNotFound
	}
	return nil
}
