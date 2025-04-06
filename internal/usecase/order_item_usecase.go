package usecase

import (
	"context"
	"errors"
	"order-service/internal/domain"
)

type orderItemUseCase struct {
	repo domain.OrderItemRepository
}

func NewOrderItemUseCase(repo domain.OrderItemRepository) domain.OrderItemUseCase {
	return &orderItemUseCase{repo: repo}
}

func (uc *orderItemUseCase) CreateOrderItem(ctx context.Context, item *domain.OrderItem) error {
	if item.Quantity <= 0 {
		return errors.New("quantity must be positive")
	}
	return uc.repo.CreateOrderItem(ctx, item)
}

func (uc *orderItemUseCase) GetOrderItemsByOrderID(ctx context.Context, orderID uint) ([]domain.OrderItem, error) {
	return uc.repo.GetOrderItemsByOrderID(ctx, orderID)
}

func (uc *orderItemUseCase) UpdateOrderItem(ctx context.Context, item *domain.OrderItem) error {
	if item.Quantity <= 0 {
		return errors.New("quantity must be positive")
	}
	return uc.repo.UpdateOrderItem(ctx, item)
}

func (uc *orderItemUseCase) DeleteOrderItem(ctx context.Context, id uint) error {
	return uc.repo.DeleteOrderItem(ctx, id)
}
