package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDataBase() *gorm.DB {
	database, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&User{})
	database.AutoMigrate(&Cart{})
	database.AutoMigrate(&Order{})
	database.AutoMigrate(&Item{})
	database.AutoMigrate(&CartItem{})

	DB = database
	return DB
}

// Using this function to get a connection, you can create your connection pool here.
func GetDb() *gorm.DB {
	return DB
}
