package services

import (
	"ecommercestorebackend/models"
)

// You could input the conditions and it will return an User in database with error info.
// 	userModel, err := FindOneUser(&User{Username: "username0"})
func FindOneUser(condition interface{}) (models.User, error) {
	database := models.GetDb()
	var user models.User

	err := database.Where(condition).First(&user).Error
	return user, err
}

// You could update properties of an User to database returning with error info.
func UpdateUser(user models.User, data interface{}) error {
	database := models.GetDb()
	err := database.Model(user).Update(data).Error
	return err
}
