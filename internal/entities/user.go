package entities

import (
	"time"
)

// User represents a user in the system
type User struct {
	UserID      uint64    `json:"user_id" db:"user_id" gorm:"column:user_id;primaryKey;autoIncrement"`
	Username    string    `json:"username" db:"username" gorm:"column:username;not null;unique;size:50"`
	Email       string    `json:"email" db:"email" gorm:"column:email;not null;unique;size:100"`
	Password    string    `json:"-" db:"password_hash" gorm:"column:password_hash;not null"` // Don't expose password in JSON
	FirstName   *string   `json:"first_name,omitempty" db:"first_name" gorm:"column:first_name;size:50"`
	LastName    *string   `json:"last_name,omitempty" db:"last_name" gorm:"column:last_name;size:50"`
	PhoneNumber *string   `json:"phone_number,omitempty" db:"phone_number" gorm:"column:phone_number;size:20"`
	IsActive    bool      `json:"is_active" db:"is_active" gorm:"column:is_active;default:true"`
	CreatedAt   time.Time `json:"created_at" db:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at" gorm:"column:updated_at"`
}

// CreateUserRequest represents the request to create a new user
type CreateUserRequest struct {
	Username    string `json:"username" binding:"required,min=3,max=50"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=6"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}
