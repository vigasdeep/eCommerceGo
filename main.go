package main

import (
	"ecommerce-backend/config"
	"ecommerce-backend/router"
	"log"

	"github.com/joho/godotenv"
)

// @title           Swagger Example API
// @version         0.1
// @description     This is a ecommerce backend built in Go.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  vigasdeep@gmail.com

// @license.name  MIT

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    config.InitDB()
    InitializeDatabase()

	r := router.SetupRouter()

	// Start the server
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
