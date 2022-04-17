package model

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID   		string   `gorm:"default:uuid_generate_v4()"`
	UserID   	string   `gorm:"default:uuid_generate_v4()"`
	OrderItems  []OrderItem
}

func (Order) TableName() string {
	return "Order"
}
