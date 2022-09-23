package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	rHandler "github.com/shukubota/go-api-template/serverless-template/registerDeviceToken/handler"
	"log"
)

func main() {
	ctx := context.Background()

	r, err := rHandler.Handler(ctx, events.APIGatewayProxyRequest{Body: "{\"fcm_device_token\": \"hoge\"}"})
	fmt.Println(r)
	if err != nil {
		log.Fatalln("handler error")
	}
}
