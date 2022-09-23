package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	handler "github.com/shukubota/go-api-template/serverless-template/registerDeviceToken/handler"
)

func main() {
	lambda.Start(handler.Handler)
}
