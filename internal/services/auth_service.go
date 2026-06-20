package services

import (
	"errors"
	"time"

	"imobiliaria-api/internal/models"
	"imobiliaria-api/internal/repositories"
	"imobiliaria-api/internal/utils"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo         repositories.UserRepository
	RefreshTokenRepo repositories.RefreshTokenRepository
}

func (s *AuthService) Login(
	email string,
	password string,
) (string, string, error) {

	user, err := s.UserRepo.FindByEmail(email)

	if err != nil {
		return "", "", errors.New("credenciais inválidas")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(password),
	)

	if err != nil {
		return "", "", errors.New("credenciais inválidas")
	}

	accessToken, err := utils.GenerateAccessToken(
		user.ID.String(),
		user.Role,
	)

	if err != nil {
		return "", "", err
	}

	refreshToken, err := utils.GenerateRefreshToken(
		user.ID.String(),
	)

	if err != nil {
		return "", "", err
	}

	refreshModel := models.RefreshToken{
		ID:         uuid.New(),
		UserID:     user.ID,
		TokenHash:  utils.HashToken(refreshToken),
		ExpiresAt:  time.Now().Add(30 * 24 * time.Hour),
		Revoked:    false,
	}

	err = s.RefreshTokenRepo.Create(&refreshModel)

	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}