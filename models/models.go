package models_practice

import (
	"fmt"
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

func PutFoodData(c *gin.Context) {
	id := c.Param("id")
	var food food_list

	if result := DB.First(&food, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	if err := c.BindJSON(&food); err != nil {
		return
	}

	fmt.Println(food)

	if result := DB.Save(&food); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, food)
}

// DeleteFoodData deletes an existing Food from the DB.
func DeleteFoodData(c *gin.Context) {
	id := c.Param("id")
	var food food_list

	if result := DB.First(&food, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	if result := DB.Delete(&food); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Food deleted successfully",
	})
}

// PatchFoodPrice updates an existing Food's price in the DB.
func PatchFoodPrice(c *gin.Context) {
	id := c.Param("id")
	var food food_list

	if result := DB.First(&food, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	var newPrice struct {
		Price float32 `json:"price"`
	}

	if err := c.BindJSON(&newPrice); err != nil {
		return
	}

	food.Price = newPrice.Price

	if result := DB.Save(&food); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, food)
}

// PatchMarket updates a Change Food's market in the DB.
func PatchMarket(c *gin.Context) {
	id := c.Param("id")
	makretID := c.Param("newMarketID")

	var food food_list

	if result := DB.First(&food, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	food.MarketID = int8(makretID[0])

	if result := DB.Save(&food); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, food)
}