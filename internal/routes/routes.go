package routes

import (
	"imobiliaria-api/internal/handlers"
	"imobiliaria-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {

	api := r.Group("/api/v1")
	categories := api.Group("/categories")
categories.Use(middleware.AuthMiddleware())
{
	categories.POST("", handlers.CreateCategory)
	categories.GET("", handlers.GetCategories)
	categories.GET("/:id", handlers.GetCategory)
	categories.PUT("/:id", handlers.UpdateCategory)
	categories.DELETE("/:id", handlers.DeleteCategory)
}

images := api.Group("/images")
images.Use(middleware.AuthMiddleware())
{
	images.POST("", handlers.UploadImage)
	images.GET("/property/:propertyID", handlers.GetPropertyImages)
	images.DELETE("/:id", handlers.DeleteImage)
}

	// AUTH
	auth := api.Group("/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
	       //auth.GET("/me", middleware.AuthMiddleware(), handlers.Me)
	}

	// PROPERTIES
	properties := api.Group("/properties")
	properties.Use(middleware.AuthMiddleware())
	{
		properties.POST("", handlers.CreateProperty)
		properties.GET("", handlers.GetProperties)
		properties.GET("/search", handlers.SearchProperties)
		properties.GET("/:id", handlers.GetProperty)
		properties.DELETE("/:id", handlers.DeleteProperty)
		properties.PUT("/:id", handlers.UpdateProperty)
		
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API Imobiliária Online",
		})
	})
}