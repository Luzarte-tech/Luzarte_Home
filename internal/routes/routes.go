package routes

import (
	"imobiliaria-api/internal/handlers"
	"imobiliaria-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {

	api := r.Group("/api/v1")

	auth := api.Group("/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)

		auth.GET(
			"/me",
			middleware.AuthMiddleware(),
			handlers.Me,
		)
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API Imobiliaria Online",
		})
	})
}