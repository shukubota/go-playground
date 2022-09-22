package main

import (
	"example/hello/serverless-template/searchContent/handler"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler.Handler)
}
