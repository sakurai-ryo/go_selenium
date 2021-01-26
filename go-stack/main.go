package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sclevine/agouti"
)

func main() {
	lambda.Start(handler)
}
func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("ctx: ", ctx)
	log.Println("request: ", request)

	title, err := scraping()
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "failed",
			StatusCode: 400,
		}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("get Title: %s", title),
		StatusCode: 200,
	}, nil
}

func scraping() (string, error) {
	driver := agouti.ChromeDriver()
	defer driver.Stop()

	if err := driver.Start(); err != nil {
		return "", fmt.Errorf("failed to start driver: %s", err)
	}

	page, err := driver.NewPage()
	log.Println("new Pages: ", page)
	if err != nil {
		return "", fmt.Errorf("failed to create NewPages: %s", err)
	}

	page.Navigate("https://www.google.com/")
	return page.Title()
}
