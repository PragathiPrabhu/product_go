package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID string
	Item string
	Price  string
}