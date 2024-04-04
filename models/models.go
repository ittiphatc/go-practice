package models_practice

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type food_list struct {
	FoodID   int8    `json:"id" gorm:"primary_key"`
	MarketID int8    `json:"market_id"`
	Food     string  `json:"food_name"`
	Price    float32 `json:"price"`
}

// Food adds an album from JSON received in the request body.
func PostFood(c *gin.Context) {
	var newFood food_list

	// Call BindJSON to bind the received JSON to
	// newFood.
	if err := c.BindJSON(&newFood); err != nil {
		return
	}

	// Add Food
	if result := DB.Create(&newFood); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, newFood)
}

// getFoodByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetFoodByID(c *gin.Context) {
	id := c.Param("id")
	var food food_list

	if result := DB.First(&food, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, food)
}

// getFood responds with the list of all albums as JSON.
func GetFood(c *gin.Context) {
	// get Food
	var food []food_list

	if result := DB.Find(&food); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, food)
}

func GetFoodinMarket(c *gin.Context) {
	id := c.Param("id")
	var food []food_list

	if result := DB.Where("market_id = ?", id).Find(&food); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, food)
}
