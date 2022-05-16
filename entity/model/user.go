package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string  `json:"name" bindding:"required"`
	Email string  `json:"email" bindding:"required"`
	Password string  `json:"password" bindding:"required"`
	Token string  `json:"token" bindding:"required"`
}
