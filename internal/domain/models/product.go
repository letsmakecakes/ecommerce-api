package models

import "time"

// Product represents a product in the e-commerce platform
type Product struct {
	ID          uint      `json:"id" gorm:"primaryKey"`             // Primary key for the product
	Name        string    `json:"name" gorm:"not null"`             // Name of the product, cannot be null
	Description string    `json:"description"`                      // Optional description of the product
	Price       float64   `json:"price" gorm:"not null"`            // Price of the product, cannot be null
	Stock       int       `json:"stock" gorm:"not null"`            // Quantity of stock, cannot be null
	ImageURL    string    `json:"image_url"`                        // Optional URL for the product image
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"` // Automatically set timestamp for record creation
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"` // Automatically set timestamp for record update
}
