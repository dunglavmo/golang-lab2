package models

import (
	"time"

	"github.com/google/uuid"
)

type Photo struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	Name      string    `gorm:"uniqueIndex;not null" json:"title,omitempty"`
	Link      string    `gorm:"not null" json:"link,omitempty"`
	User      uuid.UUID `gorm:"not null" json:"user,omitempty"`
	Album     uuid.UUID `gorm:"not null" json:"album,omitempty"`
	CreatedAt time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at,omitempty"`
}

type CreatePhotoRequest struct {
	Name      string    `json:"name"  binding:"required"`
	Link      string    `json:"link" binding:"required"`
	User      string    `json:"user,omitempty"`
	Album     string    `json:"album,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type UpdatePhoto struct {
	Name      string    `json:"name,omitempty"`
	Link      string    `json:"link,omitempty"`
	User      string    `json:"user,omitempty"`
	Album     string    `json:"album,omitempty"`
	CreateAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
