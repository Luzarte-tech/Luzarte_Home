package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Property struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	OwnerID uuid.UUID `gorm:"type:uuid;not null"`

	CategoryID uuid.UUID `gorm:"type:uuid;not null"`

	Title string `gorm:"size:255;not null"`

	Description string

	TransactionType string

	Price float64

	Bedrooms int

	Bathrooms int

	GarageSpaces int

	Area float64

	Address string

	City string

	Province string

	Latitude float64

	Longitude float64

	Status string `gorm:"default:pending"`

	Featured bool `gorm:"default:false"`

	Published bool `gorm:"default:true"`

	CreatedAt time.Time

	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

}
