package models

import "time"

// CartItem represents an item in the shopping cart.
type CartItem struct {
	ID        uint    `json:"id" gorm:"primaryKey"`                // Primary key for the cart item
	CartID    uint    `json:"cart_id"`                             // Foreign key to the cart
	ProductID uint    `json:"product_id"`                          // Foreign key to the product
	Quantity  int     `json:"quantity"`                            // Quantity of the product
	Product   Product `json:"product" gorm:"foreignKey:ProductID"` // The product associated with the cart
}

// Cart represents a shopping cart.
type Cart struct {
	ID        uint       `json:"id" gorm:"primaryKey"`             // Primary key for the cart
	UserID    uint       `json:"user_id"`                          // Foreign key to the user
	Items     []CartItem `json:"items" gorm:"foreignKey:CartID"`   // List of items in the cart
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"` // Automatically set timestamp for record creation
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime"` // Automatically set timestamp for record update
}
