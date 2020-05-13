package main

import (
	"fmt"

	"example.com/config"
	"example.com/models"
	"example.com/routes"
	"github.com/jinzhu/gorm"
)

var err error

func main() {

	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))

	if err != nil {
		fmt.Println("statuse: ", err)
	}

	defer config.DB.Close()
	config.DB.AutoMigrate(&models.Todo{})

	r := routes.SetupRouter()
	// running
	r.Run()
}
