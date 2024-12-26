package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/product", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "getall",
		})
	})

	r.Run(":8083") 
}
