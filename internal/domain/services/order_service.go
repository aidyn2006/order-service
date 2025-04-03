package services

import (
	"order-service/models"
	"order-service/repositories"
)

func CreateOrderService(order *models.Order) error {
	return repositories.CreateOrder(order)
}
