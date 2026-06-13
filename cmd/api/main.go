package main

import (
	"imobiliaria-api/internal/database"
	"imobiliaria-api/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	database.Connect()

	r := gin.Default()

	routes.Setup(r)

	r.Run(":8080")
}
