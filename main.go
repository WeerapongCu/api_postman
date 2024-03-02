package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/PRINTME", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	r.GET(("/SHOWPARAM/:TEXT"), func(c *gin.Context) {
		text := c.Param("TEXT")
		c.JSON(200, gin.H{
			"message": text,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}