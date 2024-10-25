package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name  string  `json:"name" example:"Laptop"`
	Price float64 `json:"price" example:"1299.99"`
}
