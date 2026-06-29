package repositories

import (
	"imobiliaria-api/internal/database"
	"imobiliaria-api/internal/models"
)

type ImageRepository struct{}

func (r *ImageRepository) Create(image *models.Image) error {
	return database.DB.Create(image).Error
}

func (r *ImageRepository) FindByProperty(propertyID string) ([]models.Image, error) {

	var images []models.Image

	err := database.DB.
		Where("property_id = ?", propertyID).
		Find(&images).Error

	return images, err
}

func (r *ImageRepository) FindByID(id string) (*models.Image, error) {

	var image models.Image

	err := database.DB.
		Where("id = ?", id).
		First(&image).Error

	if err != nil {
		return nil, err
	}

	return &image, nil
}

func (r *ImageRepository) Update(image *models.Image) error {
	return database.DB.Save(image).Error
}

func (r *ImageRepository) Delete(id string) error {
	return database.DB.Delete(&models.Image{}, "id = ?", id).Error
}