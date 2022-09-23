package register_device_token_handler_test

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	handler "github.com/shukubota/go-api-template/serverless-template/registerDeviceToken/handler"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Register_Device_Token_Handler(t *testing.T) {
	ctx := context.Background()
	testCases := []struct {
		requestBody   string
		expected      events.APIGatewayProxyResponse
		expectedError error
	}{
		{
			requestBody: "{\"fcm_device_token\": \"hoge\"}",
			expected: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Body:       "OK",
			},
		},
		{
			requestBody: "{\"fcm_device_token: \"hoge\"}",
			expected: events.APIGatewayProxyResponse{
				StatusCode: 400,
				Body:       "json unmarshal fail",
			},
		},
	}

	for i, tc := range testCases {
		fmt.Println(i)
		data := events.APIGatewayProxyRequest{
			Body: tc.requestBody,
		}
		r, _ := handler.Handler(ctx, data)
		// errは判定しない
		assert.Equal(t, tc.expected, r)
	}
}
