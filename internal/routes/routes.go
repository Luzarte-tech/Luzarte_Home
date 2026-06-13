package routes

import (
	"imobiliaria-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {

	api := r.Group("/api/v1")

	auth := api.Group("/auth")
	{
		auth.POST("/register", handlers.Register)
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API Imobiliaria Online",
		})
	})
}