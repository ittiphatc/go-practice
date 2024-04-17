package main

import (
	models_practice "go-project/models"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)
func main() {
	// Connect to database
	models_practice.InitialDB()

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	}))
	router.GET("/food", models_practice.GetFood)
	router.POST("/food", models_practice.PostFood)
	router.GET("/food/:id", models_practice.GetFoodByID)
	router.GET("/market/:id", models_practice.GetFoodinMarket)

	router.Run("localhost:8080")

}
