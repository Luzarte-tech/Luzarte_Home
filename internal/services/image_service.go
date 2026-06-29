package services

import (
	"imobiliaria-api/internal/models"
	"imobiliaria-api/internal/repositories"
)

type ImageService struct {
	ImageRepo repositories.ImageRepository
}

func (s *ImageService) Create(image *models.Image) error {
	return s.ImageRepo.Create(image)
}

func (s *ImageService) FindByProperty(propertyID string) ([]models.Image, error) {
	return s.ImageRepo.FindByProperty(propertyID)
}

func (s *ImageService) FindByID(id string) (*models.Image, error) {
	return s.ImageRepo.FindByID(id)
}

func (s *ImageService) Update(image *models.Image) error {
	return s.ImageRepo.Update(image)
}

func (s *ImageService) Delete(id string) error {
	return s.ImageRepo.Delete(id)
}