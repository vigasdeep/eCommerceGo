package handlers

import (
	"ecommerce-backend/config"
	"ecommerce-backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetOrderItems(c *gin.Context) {
    orderID, _ := strconv.Atoi(c.Param("order_id"))
    var orderItems []models.OrderItem

    if result := config.DB.Where("order_id = ?", orderID).Find(&orderItems); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusOK, orderItems)
}
func CreateOrderItem(c *gin.Context) {
    var orderItem models.OrderItem
    orderID, _ := strconv.Atoi(c.Param("order_id")) 

    if err := c.ShouldBindJSON(&orderItem); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    orderItem.OrderID = uint(orderID)

    if result := config.DB.Create(&orderItem); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    if err := RecalculateOrderTotal(config.DB, orderItem.OrderID); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, orderItem)
}
func UpdateOrderItem(c *gin.Context) {
    var updatedData models.OrderItem
    product_id, _ := strconv.Atoi(c.Param("product_id")) 

    if err := c.ShouldBindJSON(&updatedData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if result := config.DB.Model(&models.OrderItem{}).Where("id = ?", product_id).Updates(updatedData); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    var orderItem models.OrderItem
    config.DB.First(&orderItem, product_id)

    if err := RecalculateOrderTotal(config.DB, orderItem.OrderID); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, updatedData)
}
func DeleteOrderItem(c *gin.Context) {
    itemID, _ := strconv.Atoi(c.Param("item_id"))

    var orderItem models.OrderItem
    if result := config.DB.First(&orderItem, itemID); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Order item not found"})
        return
    }

    if result := config.DB.Delete(&models.OrderItem{}, itemID); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    if err := RecalculateOrderTotal(config.DB, orderItem.OrderID); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Order item deleted successfully"})
}