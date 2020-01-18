package main

import (
	"github.com/gin-gonic/gin"
)

// /hello endpoint
func homeLink(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}

// /greet/{name} endpoint
func nameLink(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, " + c.Param("name"),
	})
}

func main() {
	// init router
	router := gin.Default()
	// register endpoints
	router.GET("/hello", homeLink)
	router.GET("/greet/:name", nameLink)
	// print server status
	router.Run(":8088")
}
