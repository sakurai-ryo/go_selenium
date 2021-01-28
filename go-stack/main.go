package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/chromedp"
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
	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	// navigate to a page, wait for an element, click
	var res string
	err := chromedp.Run(ctx,
		emulation.SetUserAgentOverride("WebScraper 1.0"),
		chromedp.Navigate(`https://github.com`),
		// wait for footer element is visible (ie, page is loaded)
		// chromedp.ScrollIntoView(`footer`),
		chromedp.WaitVisible(`footer < div`),
		chromedp.Text(`h1`, &res, chromedp.NodeVisible, chromedp.ByQuery),
	)
	if err != nil {
		return "", err
	}
	return res, nil
}

/*
func scraping() (string, error) {
	driver := agouti.ChromeDriver(agouti.ChromeOptions(
		"binary", "../opt/headless-chromium",
	))
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
*/
