package handler

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
)

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("----------post line handler")
	fmt.Println(request)
	// ここに処理
	return events.APIGatewayProxyResponse{}, nil
}
