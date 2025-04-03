package repositories

import (
	"order-service/database"
	"order-service/models"
)

func CreateOrder(order *models.Order) error {
	return database.DB.Create(order).Error
}
