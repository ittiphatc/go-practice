package models_practice

import(
	"os"
	"log"
	"fmt"

	"github.com/joho/godotenv" // import godotenv
	"gorm.io/driver/mysql"     // import gorm mysql
	"gorm.io/gorm"             // import gorm
)

var DB *gorm.DB

func InitialDB(){
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	// PostgreSQL connected
	// dsn := fmt.Sprintf(
	// 	"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	DB = database

	DB.AutoMigrate(&food_list{})
}