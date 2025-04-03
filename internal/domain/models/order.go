package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID uint    `json:"user_id"`
	Total  float64 `json:"total"`
	Status string  `json:"status"`
}
