package models


import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	Name string `gorm:"not null"`

	Email string `gorm:"unique;not null"`

	Phone string

	PasswordHash string `gorm:"not null"`

	Role string `gorm:"default:client"`

	Status string `gorm:"default:active"`

	CreatedAt time.Time
	UpdatedAt time.Time
	RoleID uuid.UUID `gorm:"type:uuid"`
}