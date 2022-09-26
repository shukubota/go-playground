package handler

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/aws/aws-lambda-go/events"
	"log"
	"net/http"
)

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("----------sssss")
	// ここにscraping処理

	res, err := http.Get("https://qiita.com/timeline")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
		return events.APIGatewayProxyResponse{}, err
	}

	fmt.Println(res)
	fmt.Println(err)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("===========doc")
	fmt.Println(doc)

	return events.APIGatewayProxyResponse{}, nil
}
