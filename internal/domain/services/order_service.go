package services

import (
	"order-service/internal/domain/models"
	"order-service/internal/repository"
)

func CreateOrderService(order *models.Order) error {
	return repository.CreateOrder(order)
}
