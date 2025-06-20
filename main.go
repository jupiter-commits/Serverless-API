package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

func init() {
	fmt.Println("Starting...")
	r := gin.Default()
	r.GET(
		"/api",
		func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Welcome"})
		})
	ginLambda = ginadapter.New(r)
}

func HandleRequest(ctx context.Context, req events.ApiGatewayProxyRequest) (events.ApiGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	log.Println("App started")
	lambda.Start(HandleRequest)
}
