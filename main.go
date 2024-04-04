package main

import (
	"github.com/gin-gonic/gin"
	models "go-project/models"
)

func main() {

	models.initialDB()
	
	router := gin.Default()
	router.GET("/food", models.getFood())
	router.POST("/food", models.postFood())
	router.GET("/food/:id", models.getFoodByID())
	router.GET("/market/:id", models.getFoodinMarket())

	router.Run("localhost:8080")

}
