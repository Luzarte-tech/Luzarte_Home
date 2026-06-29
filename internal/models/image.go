package models

import (
	"time"

	"github.com/google/uuid"
)

type Image struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	PropertyID uuid.UUID `gorm:"type:uuid;not null"`

	FileName string `gorm:"size:255;not null"`

	FilePath string `gorm:"size:500;not null"`

	IsPrimary bool `gorm:"default:false"`

	CreatedAt time.Time
}