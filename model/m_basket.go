package model

import (
	"gorm.io/gorm"
)

type Basket struct {
	gorm.Model
	ID   			string   `gorm:"default:uuid_generate_v4()"`
	UserID  		string   `gorm:"default:uuid_generate_v4()"`
	OrderItems  	[]OrderItem
}

func (Basket) TableName() string {
	return "Basket"
}