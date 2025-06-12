package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting...")
	r := gin.Default()
	r.GET(
		"/api",
		func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Welcome"})
		})

	r.Run() // 8080
}
