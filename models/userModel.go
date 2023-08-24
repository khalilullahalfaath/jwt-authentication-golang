package models

import "gorm.io/gorm"

// User is the model for the users table
type User struct {
	gorm.Model
	Email string `gorm:"unique" `
	Password string 
}