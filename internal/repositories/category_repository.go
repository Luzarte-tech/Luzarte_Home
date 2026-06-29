package repositories

import (
	"imobiliaria-api/internal/database"
	"imobiliaria-api/internal/models"
)

type CategoryRepository struct{}

func (r *CategoryRepository) Create(category *models.Category) error {
	return database.DB.Create(category).Error
}

func (r *CategoryRepository) FindAll() ([]models.Category, error) {

	var categories []models.Category

	err := database.DB.
		Order("name ASC").
		Find(&categories).Error

	return categories, err
}

func (r *CategoryRepository) FindByID(id string) (*models.Category, error) {

	var category models.Category

	err := database.DB.
		Where("id = ?", id).
		First(&category).Error

	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *CategoryRepository) Update(category *models.Category) error {
	return database.DB.Save(category).Error
}

func (r *CategoryRepository) Delete(id string) error {
	return database.DB.Delete(&models.Category{}, "id = ?", id).Error
}