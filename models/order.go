package models

import (
	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	CartID uint `json:"cart_id"`
	UserID uint `json:"user_id"`
}
