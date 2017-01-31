package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"github.com/dgrijalva/jwt-go"
)

func main() {
	router := gin.Default()

	api := router.Group("/api")
	api.Use(AuthRequired())
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": 200,
			})
		})
		api.GET("/oauth/:token", func(c *gin.Context) {
			
			key := []byte(c.Param("token"))
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"token": c.Param("token"),
				"id": 0,
			})
			tokenString, err := token.SignedString(key)
			if err != nil {
				c.JSON(401, gin.H{
					"status": 401,
					"error": "Invalid token",
				})
			}
			c.JSON(200, gin.H{
				"status": 200,
				"access-token":  tokenString,
			})
		})
	}

	router.Run(":8931")
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {

		// token := c.Request.Header.Get("Auth-Token")

		// if token != "" {
		// 	c.Next()
		// } else {
		// 	AuthError(c, 401, "Auth-Token is required for access to API routes")
		// }
		c.Next()
	}
}

func AuthError(c *gin.Context, status int, err string) {
	c.JSON(status, gin.H{
		"status": status,
		"err":    err,
	})
	c.Abort()
}
