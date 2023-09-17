package models

import (
	"time"

	"github.com/google/uuid"
)

type Album struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name        string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:text"`
	CreatedAt   time.Time `gorm:"not null"`
	UpdatedAt   time.Time `gorm:"not null"`
}

type CreateAlbumRequest struct {
	Name        string    `json:"name"  binding:"required"`
	Description string    `json:"description"  binding:"required"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type UpdateAlbum struct {
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	CreateAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
