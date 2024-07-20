package main

import (
	"ecommerce-backend/config"
	"ecommerce-backend/models"
	"log"
)

func InitializeDatabase() {
	db := config.DB

	// Automatically migrate the schema
	err := db.AutoMigrate(&models.Product{}, &models.User{})

	if err != nil {
		log.Fatalf("Failed to migrate database: %v\n", err)
	}

	// Seed data (optional)
	seedProducts()
}

func seedProducts() {
	db := config.DB

	// Example of seeding data
	products := []models.Product{
		{Name: "Sample Product 1", Price: 100},
		{Name: "Sample Product 2", Price: 200},
	}

	for _, product := range products {
		db.Where(models.Product{Name: product.Name}).FirstOrCreate(&product)
	}
}
