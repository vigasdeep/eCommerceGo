package handlers

import (
	"ecommerce-backend/config"
	"ecommerce-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateOrder(c *gin.Context) {
    var order models.Order
    userID, exists := c.Get("userID")
	if !exists {
	    c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found in context"})
	    return
	}
    order.UserID = userID.(uint)
    
    if result := config.DB.Create(&order); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusOK, order)
}

func UpdateOrderStatus(c *gin.Context) {
    orderID := c.Param("id")
    var status struct {
        Status string `json:"status"`
    }
    if err := c.ShouldBindJSON(&status); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var order models.Order
    if result := config.DB.First(&order, orderID); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
        return
    }
    if err := RecalculateOrderTotal(config.DB, order.ID); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    order.Status = status.Status
    config.DB.Save(&order)

    c.JSON(http.StatusOK, order)
}

func RecalculateOrderTotal(db *gorm.DB, orderID uint) error {
    var orderItems []models.OrderItem
    if err := db.Where("order_id = ?", orderID).Find(&orderItems).Error; err != nil {
        return err
    }

    var total float64
    for _, item := range orderItems {
        total += item.Price * float64(item.Quantity)
    }

    return db.Model(&models.Order{}).Where("id = ?", orderID).Update("total", total).Error
}