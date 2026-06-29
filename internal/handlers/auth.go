package handlers


import (
	"net/http"

	"imobiliaria-api/internal/database"
	"imobiliaria-api/internal/models"
	"golang.org/x/crypto/bcrypt"
	"github.com/gin-gonic/gin"
	"imobiliaria-api/internal/services"
        "imobiliaria-api/internal/dto"

)

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func Register(c *gin.Context) {

	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	hash, _ := bcrypt.GenerateFromPassword(
	[]byte(req.Password),
	bcrypt.DefaultCost,
)

user := models.User{
	Name:         req.Name,
	Email:        req.Email,
	Phone:        req.Phone,
	PasswordHash: string(hash),
	Role:         "client",
}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, user)
}
func Login(c *gin.Context) {

	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	service := services.AuthService{}

	accessToken, refreshToken, err := service.Login(
		req.Email,
		req.Password,
	)

	if err != nil {
		c.JSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
func Me(c *gin.Context) {

	userID, _ := c.Get("user_id")
	role, _ := c.Get("role")

	c.JSON(200, gin.H{
		"user_id": userID,
		"role":    role,
	})
}
