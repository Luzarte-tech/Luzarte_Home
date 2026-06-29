package models

import (
	"time"

	"github.com/google/uuid"
)

type Role struct {
	ID   uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name string    `gorm:"size:50;unique;not null"`

	CreatedAt time.Time
}