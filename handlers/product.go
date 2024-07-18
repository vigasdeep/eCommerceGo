package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var products = []Product{
	{ID: 1, Name: "Product Pen", Price: 100},
	{ID: 2, Name: "Product Pencil", Price: 200},
}
// GetProducts godoc
// @Summary      Get Products
// @Description  Get all Products
// @Tags         accounts
// @Accept       json
// @Produce      json

// @Router       /products [get]
func GetProducts(c *gin.Context) {
	c.JSON(http.StatusOK, products)
}

func GetProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, product := range products {
		if product.ID == id {
			c.JSON(http.StatusOK, product)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
}

func CreateProduct(c *gin.Context) {
	var newProduct Product
	if err := c.BindJSON(&newProduct); err != nil {
		return
	}
	newProduct.ID = len(products) + 1
	products = append(products, newProduct)
	c.JSON(http.StatusCreated, newProduct)
}

func UpdateProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedProduct Product
	if err := c.BindJSON(&updatedProduct); err != nil {
		return
	}
	for i, product := range products {
		if product.ID == id {
			products[i] = updatedProduct
			c.JSON(http.StatusOK, updatedProduct)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
}

func DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, product := range products {
		if product.ID == id {
			products = append(products[:i], products[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
}