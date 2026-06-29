package handlers

import (
	"net/http"
	"os"
	"path/filepath"

	"imobiliaria-api/internal/models"
	"imobiliaria-api/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var imageService = services.ImageService{}

func UploadImage(c *gin.Context) {

	propertyID := c.PostForm("property_id")

	file, err := c.FormFile("image")

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "imagem obrigatória",
		})

		return
	}

	filename := uuid.New().String() + filepath.Ext(file.Filename)

	path := "./uploads/" + filename

	os.MkdirAll("./uploads", os.ModePerm)

	if err := c.SaveUploadedFile(file, path); err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	propertyUUID, _ := uuid.Parse(propertyID)

	image := models.Image{
		PropertyID: propertyUUID,
		FileName:   filename,
		FilePath:   path,
		IsPrimary:  false,
	}

	if err := imageService.Create(&image); err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, image)
}

func GetPropertyImages(c *gin.Context) {

	images, err := imageService.FindByProperty(
		c.Param("propertyID"),
	)

	if err != nil {

		c.JSON(500, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(200, images)
}

func DeleteImage(c *gin.Context) {

	image, err := imageService.FindByID(c.Param("id"))

	if err != nil {

		c.JSON(404, gin.H{
			"error": "imagem não encontrada",
		})

		return
	}

	os.Remove(image.FilePath)

	imageService.Delete(image.ID.String())

	c.JSON(200, gin.H{
		"message": "imagem removida",
	})
}