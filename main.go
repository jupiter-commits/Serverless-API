package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("Starting...")
	r := gin.Default()
	r.GET(
		"/api",
		func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Welcome"})
		})
}

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	return fmt.Sprintf("Hello, %$", name.Name), nil
}

func main() {
	lambda.Start(HandleRequest)
}
