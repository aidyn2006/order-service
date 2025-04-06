package handlers

import (
	"github.com/gin-gonic/gin"
)

func NewRouter(orderHandler *OrderHandler, orderItemHandler *OrderItemHandler) *gin.Engine {
	router := gin.Default()

	// Add middleware if needed
	// router.Use(middleware.Logger())

	api := router.Group("/api/v1")
	{
		orders := api.Group("/orders")
		{
			orders.POST("", orderHandler.CreateOrder)
			orders.GET("", orderHandler.ListOrders)
			orders.GET("/:id", orderHandler.GetOrder)
			orders.PUT("/:id/cancel", orderHandler.CancelOrder)
		}

		orderItems := api.Group("/order-items")
		{
			orderItems.POST("", orderItemHandler.CreateOrderItem)
			orderItems.PUT("/:id", orderItemHandler.UpdateOrderItem)
			orderItems.DELETE("/:id", orderItemHandler.DeleteOrderItem)
		}
	}

	return router
}
