package models

import "github.com/jinzhu/gorm"

type Cart struct {
	gorm.Model
	UserId      uint `json:"user_id"`
	IsPurchased bool `json:"is_purchased"`
}
