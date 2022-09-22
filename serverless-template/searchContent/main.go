package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/shukubota/go-api-template/serverless-template/searchContent/handler"
)

func main() {
	lambda.Start(handler.Handler)
}
