package services

import (
	"imobiliaria-api/internal/models"
	"imobiliaria-api/internal/repositories"
)

type PropertyService struct {
	PropertyRepo repositories.PropertyRepository
}

func (s *PropertyService) Create(property *models.Property) error {
	return s.PropertyRepo.Create(property)
}

func (s *PropertyService) FindAll(
	page int,
	limit int,
	sort string,
) ([]models.Property, error) {

	return s.PropertyRepo.FindAll(
		page,
		limit,
		sort,
	)
}

func (s *PropertyService) FindByID(id string) (*models.Property, error) {
	return s.PropertyRepo.FindByID(id)
}

func (s *PropertyService) Update(property *models.Property) error {
	return s.PropertyRepo.Update(property)
}

func (s *PropertyService) Delete(id string) error {
	return s.PropertyRepo.Delete(id)
}

func (s *PropertyService) Search(
	city string,
	transactionType string,
	status string,
	bedrooms int,
	bathrooms int,
	minPrice float64,
	maxPrice float64,
	page int,
	limit int,
	sort string,
) ([]models.Property, error) {

	return s.PropertyRepo.Search(
		city,
		transactionType,
		status,
		bedrooms,
		bathrooms,
		minPrice,
		maxPrice,
		page,
		limit,
		sort,
	)
}
func (s *PropertyService) IsOwner(
	propertyID string,
	userID string,
) (bool, error) {

	return s.PropertyRepo.IsOwner(
		propertyID,
		userID,
	)
}