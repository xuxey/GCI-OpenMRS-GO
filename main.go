package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// global const for name param
const nameParam string = "name"

// /hello endpoint
func helloEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}

// /greet/{name} endpoint
func nameEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, " + c.Param(nameParam),
	})
}

func main() {
	// init router
	router := gin.Default()
	// register endpoints
	router.GET("/hello", helloEndpoint)
	router.GET("/greet/:"+nameParam, nameEndpoint)
	// use port 8088 on localhost
	router.Run(":8088")
}
