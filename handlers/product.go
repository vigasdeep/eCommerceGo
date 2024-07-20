package handlers

import (
	"ecommerce-backend/config"
	"ecommerce-backend/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}
var products []models.Product


// GetProducts godoc
// @Summary      Get Products
// @Description  Get all Products
// @Tags         accounts
// @Accept       json
// @Produce      json

// @Router       /products [get]
func GetProducts(c *gin.Context) {
	var products []models.Product
    result := config.DB.Find(&products)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusOK, products)
}

func GetProduct(c *gin.Context) {
	id := c.Param("id")
    var product models.Product
    result := config.DB.First(&product, id)
	fmt.Printf("%v", result)
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }

    c.JSON(http.StatusOK, product)
}

func CreateProduct(c *gin.Context) {
	var product models.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    result := config.DB.Create(&product)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusCreated, product)
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
    var product models.Product
    result := config.DB.First(&product, id)
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }

    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    config.DB.Save(&product)
    c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
    result := config.DB.Delete(&models.Product{}, id)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
