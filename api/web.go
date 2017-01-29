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
			c.AbortWithStatus(401)
		}
	}
}
