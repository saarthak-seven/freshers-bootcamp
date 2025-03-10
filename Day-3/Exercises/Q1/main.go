package main

import (
	"example/Q1/config"
	"example/Q1/models"
	"example/Q1/routes"
	"fmt"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}
	defer config.DB.Close()

	config.DB.AutoMigrate(&models.User{})

	r := routes.SetupRouter()
	r.Run()
}
