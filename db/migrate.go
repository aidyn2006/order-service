package db

import (
	"fmt"
	"gorm.io/gorm"
	"order-service/db/migrations"
)

func RunMigrations(db *gorm.DB) {
	fmt.Println("Running database migrations...")
	if err := migrations.Migrate(db); err != nil {
		fmt.Println("Migration failed:", err)
	} else {
		fmt.Println("Migrations applied successfully!")
	}
}
