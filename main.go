package main

import (
	models_practice "go-project/models"

	"github.com/gin-gonic/gin"
)

func main() {

	models_practice.InitialDB()

	router := gin.Default()
	router.GET("/food", models_practice.GetFood)
	router.POST("/food", models_practice.PostFood)
	router.GET("/food/:id", models_practice.GetFoodByID)
	router.GET("/market/:id", models_practice.GetFoodinMarket)

	router.Run("localhost:8080")

}
