package models

import "time"

// User represents a user in the system.
type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`         // Primary key
	Email     string    `json:"email" gorm:"unique;not null"` // Unique and not null
	Password  string    `json:"-" gorm:"not null"`            // Password is not exposed in JSON and is not null
	FirstName string    `json:"first_name"`                   // Optional first name
	LastName  string    `json:"last_name"`                    // Optional last name
	Role      string    `json:"role" gorm:"default:'user'"`   // Role defaults to `user`
	CreatedAt time.Time `json:"created_at"`                   // Timestamp for record creation
	UpdatedAt time.Time `json:"updated_at"`                   // Timestamp for record update
}
