package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"order-service/internal/domain"
	"strconv"
)

type OrderItemHandler struct {
	orderItemUseCase domain.OrderItemUseCase
}

func NewOrderItemHandler(uc domain.OrderItemUseCase) *OrderItemHandler {
	return &OrderItemHandler{orderItemUseCase: uc}
}

func (h *OrderItemHandler) CreateOrderItem(c *gin.Context) {
	var item domain.OrderItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.orderItemUseCase.CreateOrderItem(c.Request.Context(), &item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, item)
}

func (h *OrderItemHandler) UpdateOrderItem(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order item ID"})
		return
	}

	var item domain.OrderItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item.ID = uint(id)

	if err := h.orderItemUseCase.UpdateOrderItem(c.Request.Context(), &item); err != nil {
		if errors.Is(err, errors.New("order item not found")) {
			c.JSON(http.StatusNotFound, gin.H{"error": "order item not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *OrderItemHandler) DeleteOrderItem(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order item ID"})
		return
	}

	if err := h.orderItemUseCase.DeleteOrderItem(c.Request.Context(), uint(id)); err != nil {
		if errors.Is(err, errors.New("order item not found")) {
			c.JSON(http.StatusNotFound, gin.H{"error": "order item not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
