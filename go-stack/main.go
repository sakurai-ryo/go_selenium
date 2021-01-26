package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}
func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("ctx: ", ctx)
	log.Println("request: ", request)

	return events.APIGatewayProxyResponse{
		Body:       "success",
		StatusCode: 200,
	}, nil
}
