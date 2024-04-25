package models_practice

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{
				"message": "Unauthorized via middleware",
			})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("there was an error")
			}
			return []byte("food_secret"), nil
		})

		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("username", claims["username"])
			c.Set("ID", claims["ID"])
			c.Set("roles", claims["roles"])
			c.Next()
		} else {
			c.JSON(401, gin.H{
				"message": "Unauthorized login",
			})
			c.Abort()
			return
		}
		if token.Claims.(jwt.MapClaims)["roles"] == CUSTOMER {
			if c.Request.URL.Path == "/market/:MarketID" {
				c.JSON(403, gin.H{
					"message": "Forbidden",
				})
				c.Abort()
				return
			}
			c.Next()
		}
		if token.Claims.(jwt.MapClaims)["roles"] == RESTAURANT {
			id := c.Param("MarketID")
			if id != token.Claims.(jwt.MapClaims)["ID"] {
				c.JSON(403, gin.H{
					"message": "Forbidden",
				})
				c.Abort()
				return
			}
			c.Next()
		}
	}
}
