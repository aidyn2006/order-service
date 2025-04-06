// @title Order Service API
// @version 1.0
// @description This is an API for managing orders
// @host localhost:8081
// @BasePath /
// @schemes http
package main

import (
	"log"
	"order-service/config"
	"order-service/db"
	"order-service/internal/delivery/http/handlers"
	"order-service/internal/repository"
	"order-service/internal/usecase"
	"order-service/pkg/inventory"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize database connection
	dbConn, err := db.NewPostgresDB(cfg.DB)
	if err != nil {
		log.Fatal(err)
	}

	// Run migrations
	db.RunMigrations(dbConn)

	// Initialize inventory client
	inventoryClient := inventory.NewClient(cfg)

	// Initialize repositories
	orderRepo := repository.NewOrderRepository(dbConn)
	orderItemRepo := repository.NewOrderItemRepository(dbConn)

	// Initialize use cases
	orderUseCase := usecase.NewOrderUseCase(orderRepo, orderItemRepo, inventoryClient)
	orderItemUseCase := usecase.NewOrderItemUseCase(orderItemRepo)

	// Initialize HTTP handlers
	orderHandler := handlers.NewOrderHandler(orderUseCase)
	orderItemHandler := handlers.NewOrderItemHandler(orderItemUseCase)

	// Setup router
	router := handlers.NewRouter(orderHandler, orderItemHandler)

	// Start server
	log.Printf("Order Service running on port %s", cfg.ServerPort)
	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal(err)
	}
}
