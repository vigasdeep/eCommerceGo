package models

import "gorm.io/gorm"

type Product struct {
    gorm.Model
    Name     string `json:"name"`
    Price    float64 `json:"price"`
    Quantity int `json:"quantity"`
    UserID   uint `json:"user_id"`  // Foreign key to User
}