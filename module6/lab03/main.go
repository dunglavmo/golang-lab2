package main

import (
	"your-module-name/config"
	"your-module-name/routes"
)

func main() {
	err := config.InitDB()
	if err != nil {
		panic("Failed to connect to database")
	}
	defer config.CloseDB()

	r := routes.SetupRouter()
	r.Run(":8080")
}
