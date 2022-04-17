package model

import (
	"gorm.io/gorm"
)

type OrderItem struct {
	gorm.Model
	ID   			string   `gorm:"default:uuid_generate_v4()"`
	OrderAmount  	int64    
	ItemID			string   `gorm:"default:uuid_generate_v4()"`
}

func (OrderItem) TableName() string {
	return "OrderItem"
}