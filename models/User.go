package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Fullname    string `binding:"required"`
	Email       string `gorm:"unique" binding:"required,email"`
	Password    string `binding:"required,min=6"`
	PhoneNumber int    `binding:"required"`
}

type LoginRequest struct {
	Email    string `binding:"required,email"`
	Password string `binding:"required,min=6"`
}
