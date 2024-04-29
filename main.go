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
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	router.POST("/login", models_practice.UserLogin)
	router.POST("/register", models_practice.UserRegister)
	router.GET("/food", models_practice.GetFood)

	router.Use(models_practice.AuthMiddleware())
	{
		router.GET("/usersession", models_practice.GetUser)
		router.POST("/market/:MarkerID", models_practice.PostFood)
		router.GET("/food/:id", models_practice.GetFoodByID)
		router.GET("/market/:id", models_practice.GetFoodinMarket)
		router.PUT("/market/:MarketID/:id", models_practice.PutFoodData)
		router.DELETE("/market/:MarketID/:id", models_practice.DeleteFoodData)
		router.PATCH("/market/:MarketID/:id", models_practice.PatchFoodPrice)
	}

	router.Run(":8080")
}
