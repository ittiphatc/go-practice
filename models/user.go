package models_practice

import (
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserID   int8   `json:"id" gorm:"primary_key"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Roles    int8   `json:"roles"`
}

const (
	CUSTOMER  = 1
	RESTAURANT = 2
	ADMIN     = 3
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}


func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func UserRegister(c *gin.Context) {
	var user User
	c.BindJSON(&user)
	user.Password, _ = HashPassword(user.Password)
	DB.Create(&user)
	c.JSON(200, gin.H{
		"message": "User has been registered",
	})
}

func UserLogin(c *gin.Context) {
	var user User
	var userDB User
	c.BindJSON(&user)
	DB.Where("username = ?", user.Username).First(&userDB)
	if user.Username == userDB.Username && CheckPasswordHash(user.Password, userDB.Password) {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": user.Username,
			"roles": userDB.Roles,
		})
		tokenString, err := token.SignedString([]byte("food_secret"))
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Error while generating token",
			})
		}
		c.JSON(200, gin.H{
			"token": tokenString,
		})
	} else {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
		})
	}
}
