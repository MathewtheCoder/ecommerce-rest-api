package models

import (
	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	CartID Cart `json:"cart_id"`
	UserID User `json:"user_id"`
}
