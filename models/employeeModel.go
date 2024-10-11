package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Position string `json:"position"`
}