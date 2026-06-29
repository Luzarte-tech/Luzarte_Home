package services

import (
	"imobiliaria-api/internal/models"
	"imobiliaria-api/internal/repositories"
)

type CategoryService struct {
	CategoryRepo repositories.CategoryRepository
}

func (s *CategoryService) Create(category *models.Category) error {
	return s.CategoryRepo.Create(category)
}

func (s *CategoryService) FindAll() ([]models.Category, error) {
	return s.CategoryRepo.FindAll()
}

func (s *CategoryService) FindByID(id string) (*models.Category, error) {
	return s.CategoryRepo.FindByID(id)
}

func (s *CategoryService) Update(category *models.Category) error {
	return s.CategoryRepo.Update(category)
}

func (s *CategoryService) Delete(id string) error {
	return s.CategoryRepo.Delete(id)
}