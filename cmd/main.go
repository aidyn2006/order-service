package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"order-service/internal/domain/services"
	"order-service/internal/repository"
	"order-service/internal/transport/http"
	"order-service/pkg/database"
)

func main() {
	db := database.InitDB()

	orderRepo := repository.NewOrderRepository(db)
	orderService := services.NewOrderService(orderRepo)
	orderHandler := http.NewOrderHandler(orderService)

	router := gin.Default()

	router.POST("/orders", orderHandler.CreateOrder)
	router.GET("/orders/:id", orderHandler.GetOrderById)
	router.PATCH("/orders/:id", orderHandler.UpdateOrderStatus)

	log.Println("Order Service is running on port 8081...")
	router.Run(":8081")
}
