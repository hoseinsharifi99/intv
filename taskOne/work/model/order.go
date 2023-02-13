package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	OrderID uint `gorm:"unique_index"`
	Price   uint
	Title   string
}
