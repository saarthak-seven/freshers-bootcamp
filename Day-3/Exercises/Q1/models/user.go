package models

import (
	"example/Q1/config"
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetAllUsers(user *[]User) error {
	if err := config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

func CreateUser(user *User) error {
	if err := config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByID(user *User, id string) error {
	if err := config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *User, id string) error {
	fmt.Println(user)
	config.DB.Save(user)
	return nil
}

func DeleteUser(user *User, id string) error {
	config.DB.Where("id = ?", id).Delete(user)
	return nil
}
