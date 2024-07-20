package models

import "gorm.io/gorm"

type Product struct {
    gorm.Model
    Name     string `json:"name"`
    Price    float64 `json:"price"`
    Quantity int `json:"quantity"`
    UserID   uint `json:"user_id"`  // Foreign key to User
}
type User struct {
    gorm.Model
    Email    string `gorm:"unique" json:"email"`
    Password string `json:"password"`
    Products []Product `gorm:"foreignKey:UserID"`
}

type Order struct {
    gorm.Model
    UserID      uint       `json:"user_id"`
    Total       float64    `json:"total"`
    Status      string     `json:"status"`
    OrderItems  []OrderItem `gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
    gorm.Model
    OrderID   uint    `json:"order_id"`
    ProductID uint    `json:"product_id"`
    Quantity  int     `json:"quantity"`
    Price     float64 `json:"price"`
}
