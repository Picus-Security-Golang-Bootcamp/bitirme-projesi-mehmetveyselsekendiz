package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID   		string   `gorm:"default:uuid_generate_v4()"`
	Email       string
	Password    string
	Basket		Basket
	Orders		[]Order
}

func (User) TableName() string {
	return "User"
}