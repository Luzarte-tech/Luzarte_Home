package handlers

import (
	"net/http"
	"strconv"
	"fmt"
	"imobiliaria-api/internal/dto"
	"imobiliaria-api/internal/models"
	"imobiliaria-api/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var propertyService = services.PropertyService{}

func CreateProperty(c *gin.Context) {

	var req dto.CreatePropertyRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ownerID, _ := c.Get("user_id")

	categoryUUID, err := uuid.Parse(req.CategoryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "category_id inválido",
		})
		return
	}

	ownerUUID, err := uuid.Parse(ownerID.(string))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "utilizador inválido",
		})
		return
	}

	property := models.Property{
		OwnerID: ownerUUID,

		CategoryID: categoryUUID,

		Title: req.Title,

		Description: req.Description,

		TransactionType: req.TransactionType,

		Price: req.Price,

		Bedrooms: req.Bedrooms,

		Bathrooms: req.Bathrooms,

		GarageSpaces: req.GarageSpaces,

		Area: req.Area,

		Address: req.Address,

		City: req.City,

		Province: req.Province,

		Latitude: req.Latitude,

		Longitude: req.Longitude,

		Featured: req.Featured,

		Published: req.Published,
	}

	err = propertyService.Create(&property)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, property)
}

func GetProperties(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	sort := c.DefaultQuery("sort", "-created")

	properties, err := propertyService.FindAll(
	page,
	limit,
	sort,
)

	if err != nil {

		c.JSON(500, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(200, properties)

}

func GetProperty(c *gin.Context) {

	property, err := propertyService.FindByID(c.Param("id"))

	if err != nil {
		c.JSON(404, gin.H{
			"error": "imóvel não encontrado",
		})
		return
	}

	c.JSON(200, property)
}

func DeleteProperty(c *gin.Context) {

	err := propertyService.Delete(c.Param("id"))
userID, _ := c.Get("user_id")
role, _ := c.Get("role")

if role != "admin" {

	ok, err := propertyService.IsOwner(
		c.Param("id"),
		userID.(string),
	)

	if err != nil {

		c.JSON(500, gin.H{
			"error": err.Error(),
		})

		return
	}

	if !ok {

		c.JSON(403, gin.H{
			"error": "sem permissão",
		})

		return
	}
}

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "imóvel removido com sucesso",
	})
}

func UpdateProperty(c *gin.Context) {

	var req dto.UpdatePropertyRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	property, err := propertyService.FindByID(c.Param("id"))
userID, _ := c.Get("user_id")
role, _ := c.Get("role")

if role != "admin" {

	ok, err := propertyService.IsOwner(
		c.Param("id"),
		userID.(string),
	)

	if err != nil {

		c.JSON(500, gin.H{
			"error": err.Error(),
		})

		return
	}

	if !ok {

		c.JSON(403, gin.H{
			"error": "sem permissão",
		})

		return
	}
}

	if err != nil {
		c.JSON(404, gin.H{
			"error": "imóvel não encontrado",
		})
		return
	}

	if req.Title != "" {
		property.Title = req.Title
	}

	if req.Description != "" {
		property.Description = req.Description
	}

	if req.TransactionType != "" {
		property.TransactionType = req.TransactionType
	}

	if req.Price > 0 {
		property.Price = req.Price
	}

	property.Bedrooms = req.Bedrooms
	property.Bathrooms = req.Bathrooms
	property.GarageSpaces = req.GarageSpaces
	property.Area = req.Area

	if req.Address != "" {
		property.Address = req.Address
	}

	if req.City != "" {
		property.City = req.City
	}

	if req.Province != "" {
		property.Province = req.Province
	}

	property.Latitude = req.Latitude
	property.Longitude = req.Longitude

	if req.Status != "" {
		property.Status = req.Status
	}

	property.Featured = req.Featured
	property.Published = req.Published

	err = propertyService.Update(property)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, property)
}
func SearchProperties(c *gin.Context) {

	city := c.Query("city")
	transactionType := c.Query("transaction_type")
	status := c.Query("status")
	sort := c.DefaultQuery("sort", "-created")

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	bedrooms, _ := strconv.Atoi(c.DefaultQuery("bedrooms", "0"))
	bathrooms, _ := strconv.Atoi(c.DefaultQuery("bathrooms", "0"))

	var minPrice float64
	var maxPrice float64

	fmt.Sscanf(c.Query("min_price"), "%f", &minPrice)
	fmt.Sscanf(c.Query("max_price"), "%f", &maxPrice)

	properties, err := propertyService.Search(
		city,
		transactionType,
		status,
		bedrooms,
		bathrooms,
		minPrice,
		maxPrice,
		page,
		limit,
		sort,
	)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, properties)
}