package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = []byte("super-secret-key")

func GenerateAccessToken(
	userID string,
	role string,
) (string, error) {

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub":  userID,
			"role": role,
			"exp":  time.Now().Add(15 * time.Minute).Unix(),
		},
	)

	return token.SignedString(SecretKey)
}

func GenerateRefreshToken(
	userID string,
) (string, error) {

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": userID,
			"exp": time.Now().Add(30 * 24 * time.Hour).Unix(),
		},
	)

	return token.SignedString(SecretKey)
}