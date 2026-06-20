package repositories

import (
	"imobiliaria-api/internal/database"
	"imobiliaria-api/internal/models"
)

type RefreshTokenRepository struct{}

func (r *RefreshTokenRepository) Create(
	token *models.RefreshToken,
) error {

	return database.DB.Create(token).Error
}

func (r *RefreshTokenRepository) FindByHash(
	hash string,
) (*models.RefreshToken, error) {

	var token models.RefreshToken

	err := database.DB.
		Where("token_hash = ?", hash).
		First(&token).Error

	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (r *RefreshTokenRepository) Revoke(
	id string,
) error {

	return database.DB.
		Model(&models.RefreshToken{}).
		Where("id = ?", id).
		Update("revoked", true).
		Error
}