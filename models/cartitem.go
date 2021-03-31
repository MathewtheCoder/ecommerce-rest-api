package models

import (
	"github.com/jinzhu/gorm"
)

type CartItem struct {
	gorm.Model
	CartID uint `json:"cart_id"`
	UserID uint `json:"user_id"`
}
