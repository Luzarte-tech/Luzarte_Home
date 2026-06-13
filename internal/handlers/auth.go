package handlers


import (
	"net/http"

	"imobiliaria-api/internal/database"
	"imobiliaria-api/internal/models"
	"golang.org/x/crypto/bcrypt"
	"github.com/gin-gonic/gin"

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
	Role:         req.Role,
}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, user)
}