package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name         string
	Email        string
	Phone        string
	PasswordHash string
	Role         string
	CreatedAt    time.Time
}