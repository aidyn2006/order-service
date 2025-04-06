package migrations

import (
	"gorm.io/gorm"
	"order-service/internal/domain"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&domain.Order{},
		&domain.OrderItem{},
	)
}
