package model

import  "gorm.io/gorm"


type Task struct {
	gorm.Model
	Body string  `json:"body" bindding:"required"`
	Completed bool  `json:"completed" bindding:"required"`
}