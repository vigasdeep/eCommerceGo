package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Define a GET endpoint
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// // Define product routes
	// r.GET("/products", getProducts)
	// r.GET("/products/:id", getProduct)
	// r.POST("/products", createProduct)
	// r.PUT("/products/:id", updateProduct)
	// r.DELETE("/products/:id", deleteProduct)

	return r
}
