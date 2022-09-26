package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/shukubota/go-api-template/serverless-template/searchContent/handler"
	"log"
)

func main() {
	req := events.APIGatewayProxyRequest{}

	r, err := handler.Handler(context.Background(), req)
	fmt.Println(r)
	if err != nil {
		log.Fatalln("handler error")
	}
}
