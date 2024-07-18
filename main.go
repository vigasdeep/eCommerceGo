package main

import "ecommerce-backend/router"

func main() {
  r := router.SetupRouter()

  // Start the server
  r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
