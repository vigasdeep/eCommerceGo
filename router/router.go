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

	// Auth routes
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
		
		auth.POST("/orders", handlers.CreateOrder)
        auth.PUT("/orders/:id", handlers.UpdateOrderStatus)

		auth.POST("/orders/:order_id/items", handlers.CreateOrderItem)
        auth.PUT("/orders/items/:item_id", handlers.UpdateOrderItem)
        auth.DELETE("/orders/items/:item_id", handlers.DeleteOrderItem)
        auth.GET("/orders/:order_id/items", handlers.GetOrderItems)

    }

	return r
}
