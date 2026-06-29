package handlers

import (

	"imobiliaria-api/internal/dto"
	"imobiliaria-api/internal/models"
	"imobiliaria-api/internal/services"

	"github.com/gin-gonic/gin"
)

var categoryService = services.CategoryService{}

func CreateCategory(c *gin.Context) {

	var req dto.CreateCategoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	category := models.Category{
		Name:        req.Name,
		Description: req.Description,
	}

	if err := categoryService.Create(&category); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(201, category)
}

func GetCategories(c *gin.Context) {

	categories, err := categoryService.FindAll()

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, categories)
}

func GetCategory(c *gin.Context) {

	category, err := categoryService.FindByID(c.Param("id"))

	if err != nil {
		c.JSON(404, gin.H{
			"error": "categoria não encontrada",
		})
		return
	}

	c.JSON(200, category)
}

func UpdateCategory(c *gin.Context) {

	var req dto.UpdateCategoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	category, err := categoryService.FindByID(c.Param("id"))

	if err != nil {
		c.JSON(404, gin.H{
			"error": "categoria não encontrada",
		})
		return
	}

	if req.Name != "" {
		category.Name = req.Name
	}

	if req.Description != "" {
		category.Description = req.Description
	}

	if err := categoryService.Update(category); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, category)
}

func DeleteCategory(c *gin.Context) {

	if err := categoryService.Delete(c.Param("id")); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "categoria removida",
	})
}