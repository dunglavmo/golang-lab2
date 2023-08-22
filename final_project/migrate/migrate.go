package main

import (
	"fmt"

	"social_app/config"
	"social_app/models"
)

func init() {
	config.LoadEnv()
	config.ConnectDB()
}

func main() {
	config.DB.AutoMigrate(&models.User{})
	fmt.Println("? Migration complete")
}
