package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName     string `json:"username" validate:"required"`
	Email        string `json:"email" validate:"required"`
	HashPassword string `json:"hash_password" validate:"required"`
}
