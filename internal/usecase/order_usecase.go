package usecase

import (
	"context"
	"errors"
	"fmt"
	"order-service/internal/domain"
	"order-service/pkg/inventory"
)

type orderUseCase struct {
	orderRepo       domain.OrderRepository
	orderItemRepo   domain.OrderItemRepository
	inventoryClient *inventory.Client
}

func NewOrderUseCase(
	orderRepo domain.OrderRepository,
	orderItemRepo domain.OrderItemRepository,
	inventoryClient *inventory.Client,
) domain.OrderUseCase {
	return &orderUseCase{
		orderRepo:       orderRepo,
		orderItemRepo:   orderItemRepo,
		inventoryClient: inventoryClient,
	}
}

func (uc *orderUseCase) CreateOrder(ctx context.Context, order *domain.Order) error {
	// Validate order items
	if len(order.Items) == 0 {
		return errors.New("order must have at least one item")
	}

	// Check product availability and set prices
	for i := range order.Items {
		item := &order.Items[i]
		product, err := uc.inventoryClient.GetProduct(ctx, item.ProductID)
		if err != nil {
			return fmt.Errorf("failed to get product %d: %v", item.ProductID, err)
		}

		if product.Stock < item.Quantity {
			return fmt.Errorf("insufficient stock for product %d", item.ProductID)
		}

		item.Price = product.Price
	}

	// Calculate total
	total, err := uc.CalculateOrderTotal(ctx, order)
	if err != nil {
		return err
	}
	order.Total = total

	// Set initial status
	order.Status = domain.StatusPending

	// Create order in transaction
	err = uc.orderRepo.CreateOrder(ctx, order)
	if err != nil {
		return fmt.Errorf("failed to create order: %v", err)
	}

	// Update inventory
	for _, item := range order.Items {
		err := uc.inventoryClient.UpdateProductStock(ctx, item.ProductID, -item.Quantity)
		if err != nil {
			// In real application, you might want to compensate here
			return fmt.Errorf("failed to update inventory for product %d: %v", item.ProductID, err)
		}
	}

	return nil
}

func (uc *orderUseCase) GetOrderByID(ctx context.Context, id uint) (*domain.Order, error) {
	return uc.orderRepo.GetOrderByID(ctx, id)
}

func (uc *orderUseCase) UpdateOrderStatus(ctx context.Context, order *domain.Order) error {
	// Validate status transition
	_, err := uc.orderRepo.GetOrderByID(ctx, order.ID)
	if err != nil {
		return err
	}

	// Add any additional validation logic for status transitions here

	return uc.orderRepo.UpdateOrder(ctx, order)
}

func (uc *orderUseCase) ListOrders(ctx context.Context, userID uint) ([]domain.Order, error) {
	return uc.orderRepo.ListOrders(ctx, userID)
}

func (uc *orderUseCase) CancelOrder(ctx context.Context, id uint) error {
	// Get order to check if it can be canceled
	order, err := uc.orderRepo.GetOrderByID(ctx, id)
	if err != nil {
		return err
	}

	if order.Status == domain.StatusShipped {
		return errors.New("cannot cancel shipped order")
	}

	// Cancel order
	err = uc.orderRepo.CancelOrder(ctx, id)
	if err != nil {
		return err
	}

	// Restore inventory
	for _, item := range order.Items {
		if err := uc.inventoryClient.UpdateProductStock(ctx, item.ProductID, item.Quantity); err != nil {
			// Log the error but don't fail the cancellation
			fmt.Printf("failed to restore inventory for product %d: %v\n", item.ProductID, err)
		}
	}

	return nil
}

func (uc *orderUseCase) CalculateOrderTotal(ctx context.Context, order *domain.Order) (float64, error) {
	var total float64
	for _, item := range order.Items {
		total += item.Price * float64(item.Quantity)
	}
	return total, nil
}
