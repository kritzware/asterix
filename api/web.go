package main

import (
	"gopkg.in/gin-gonic/gin.v1"
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
	}

	router.Run(":8931")
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.Request.Header.Get("Auth-Token")

		if token != "" {
			c.Next()
		} else {
			AuthError(c, 401, "Auth-Token is required for access to API routes")
		}
	}
}

func AuthError(c *gin.Context, status int, err string) {
	c.JSON(status, gin.H{
		"status": status,
		"err":    err,
	})
	c.Abort()
}
