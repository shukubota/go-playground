package handler

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
)

type Response struct {
	Message string `json:"message"`
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("----------post line handler")
	fmt.Println(request)
	// ここに処理
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "OK",
	}, nil
}
