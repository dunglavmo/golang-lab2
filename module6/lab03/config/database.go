package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // SQLite driver
)

var DB *gorm.DB

func InitDB() error {
	var err error
	DB, err = gorm.Open("sqlite3", "todo.db")
	if err != nil {
		return err
	}
	DB.AutoMigrate(&models.Todo{})
	return nil
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
