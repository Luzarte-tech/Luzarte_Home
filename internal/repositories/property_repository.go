package repositories

import (
	"imobiliaria-api/internal/database"
	"imobiliaria-api/internal/models"
)

type PropertyRepository struct{}

func (r *PropertyRepository) Create(property *models.Property) error {
	return database.DB.Create(property).Error
}

func (r *PropertyRepository) FindAll(
	page int,
	limit int,
	sort string,
) ([]models.Property, error) {

	var properties []models.Property

	offset := (page - 1) * limit

	query := database.DB.
		Limit(limit).
		Offset(offset)

	switch sort {

	case "price":
		query = query.Order("price ASC")

	case "-price":
		query = query.Order("price DESC")

	default:
		query = query.Order("created_at DESC")
	}

	err := query.Find(&properties).Error

	return properties, err
}
func (r *PropertyRepository) FindByID(id string) (*models.Property, error) {

	var property models.Property

	err := database.DB.
		Where("id = ?", id).
		First(&property).Error

	if err != nil {
		return nil, err
	}

	return &property, nil
}

func (r *PropertyRepository) Update(property *models.Property) error {
	return database.DB.Save(property).Error
}

func (r *PropertyRepository) Delete(id string) error {
	return database.DB.Delete(&models.Property{}, "id = ?", id).Error
}
func (r *PropertyRepository) Search(
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

    var properties []models.Property

    query := database.DB.Model(&models.Property{})

    if city != "" {
        query = query.Where("LOWER(city)=LOWER(?)", city)
    }

    if transactionType != "" {
        query = query.Where("transaction_type = ?", transactionType)
    }

    if status != "" {
        query = query.Where("status = ?", status)
    }

    if bedrooms > 0 {
        query = query.Where("bedrooms = ?", bedrooms)
    }

    if bathrooms > 0 {
        query = query.Where("bathrooms = ?", bathrooms)
    }

    if minPrice > 0 {
        query = query.Where("price >= ?", minPrice)
    }

    if maxPrice > 0 {
        query = query.Where("price <= ?", maxPrice)
    }

    switch sort {

    case "price":
        query = query.Order("price asc")

    case "-price":
        query = query.Order("price desc")

    default:
        query = query.Order("created_at desc")
    }

    if page <= 0 {
        page = 1
    }

    if limit <= 0 {
        limit = 10
    }

    offset := (page - 1) * limit

    err := query.
        Offset(offset).
        Limit(limit).
        Find(&properties).Error

    return properties, err
}
func (r *PropertyRepository) IsOwner(
	propertyID string,
	userID string,
) (bool, error) {

	var count int64

	err := database.DB.
		Model(&models.Property{}).
		Where("id = ? AND owner_id = ?", propertyID, userID).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}