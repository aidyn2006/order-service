package repository

import (
	"order-service/database"
	"order-service/internal/domain/models"
)

func CreateOrder(order *models.Order) error {
	return database.DB.Create(order).Error
}
