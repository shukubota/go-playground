package register_device_token_handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
)

type RequestData struct {
	FcmDeviceToken  string `json:"fcm_device_token"`
	ApnsDeviceToken string `json:"apns_device_token"`
}

type ResponseData struct {
	FcmDeviceToken  string `json:"fcm_device_token"`
	ApnsDeviceToken string `json:"apns_device_token"`
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println(ctx)
	fmt.Println(request.Body)
	var data RequestData
	if err := json.Unmarshal([]byte(request.Body), &data); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "json unmarshal fail",
		}, err
	}

	_, err := register(&data)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "register fail",
		}, err
	}
	return events.APIGatewayProxyResponse{
		Body:       "OK",
		StatusCode: 200,
	}, nil
}

func register(data *RequestData) (bool, error) {
	// デバイストークン登録処理
	fmt.Printf("FcmDeviceToken: %v\n", data.FcmDeviceToken)
	fmt.Printf("ApnsDeviceToken: %v\n", data.ApnsDeviceToken)
	return true, nil
}
