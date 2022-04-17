package model

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID   		string   `gorm:"default:uuid_generate_v4()"`
	Name        string
	Items       []Item
}

func (Category) TableName() string {
	return "Category"
}