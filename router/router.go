package router

import (
	"ecommerce-backend/handlers"
	"ecommerce-backend/handlers/middleware"
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

	r.POST("/register", handlers.Register)
    r.POST("/login", handlers.Login)

    // Protected routes
    auth := r.Group("/")
    auth.Use(middleware.AuthMiddleware())
    {
        auth.GET("/products", handlers.GetProducts)
        auth.GET("/products/:id", handlers.GetProduct)
        auth.POST("/products", handlers.CreateProduct)
        auth.PUT("/products/:id", handlers.UpdateProduct)
        auth.DELETE("/products/:id", handlers.DeleteProduct)
    }
	// Import the handlers package

	return r
}
