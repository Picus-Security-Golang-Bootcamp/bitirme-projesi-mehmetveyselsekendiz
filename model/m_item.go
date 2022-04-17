package model

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	ID   		string   `gorm:"default:uuid_generate_v4()"`
	Name        string
	Amount      int64
	Price       float64
	CategoryID  string   `gorm:"default:uuid_generate_v4()"`
	Category	Category
}

func (Item) TableName() string {
	return "Item"
}