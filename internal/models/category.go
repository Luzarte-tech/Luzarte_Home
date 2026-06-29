package models

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	Name string `gorm:"size:100;unique;not null"`

	Description string

	CreatedAt time.Time

	UpdatedAt time.Time
}