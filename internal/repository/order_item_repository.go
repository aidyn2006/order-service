package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"order-service/internal/domain"
)

type orderItemRepository struct {
	db *gorm.DB
}

func NewOrderItemRepository(db *gorm.DB) domain.OrderItemRepository {
	return &orderItemRepository{db: db}
}

func (r *orderItemRepository) CreateOrderItem(ctx context.Context, item *domain.OrderItem) error {
	return r.db.WithContext(ctx).Create(item).Error
}

func (r *orderItemRepository) GetOrderItemsByOrderID(ctx context.Context, orderID uint) ([]domain.OrderItem, error) {
	var items []domain.OrderItem
	err := r.db.WithContext(ctx).Where("order_id = ?", orderID).Find(&items).Error
	return items, err
}

func (r *orderItemRepository) UpdateOrderItem(ctx context.Context, item *domain.OrderItem) error {
	result := r.db.WithContext(ctx).Model(&domain.OrderItem{}).Where("id = ?", item.ID).Updates(item)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("order item not found")
	}
	return nil
}

func (r *orderItemRepository) DeleteOrderItem(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).Delete(&domain.OrderItem{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("order item not found")
	}
	return nil
}
