package models

import "time"

type Book struct {
	ID        int       `gorm:"primary_key"`
	Name      string    `gorm:"column:name"`
	Author    string    `gorm:"column:author"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
