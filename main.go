package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/greet", func(c *gin.Context) {
		name := c.Query("name")
		if name == "" {
			name = "Guest"
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": "Hello, " + name + ". Greetings..!",
		})
	})

	router.Run(":8080")
}
