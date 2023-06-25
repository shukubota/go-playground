package handler

import (
	example "connect-example/gen/greet/v1"
	"connect-example/gen/greet/v1/exampleconnect"
	"context"
	"github.com/bufbuild/connect-go"
	"log"
)

type greetHandler struct{}

func NewGreetHandler() exampleconnect.GreetServiceHandler {
	return &greetHandler{}
}

func (s *greetHandler) Greet(ctx context.Context, req *connect.Request[example.GreetRequest]) (*connect.Response[example.GreetResponse], error) {
	log.Println("RequestHeaders: ", req.Header())

	res := connect.NewResponse(&example.GreetResponse{
		Greeting: "Hello " + req.Msg.Name,
	})
	return res, nil
}

func (s *greetHandler) Echo(ctx context.Context, c *connect.Request[example.StringMessage]) (*connect.Response[example.StringMessage], error) {
	//TODO implement me
	panic("implement me")
}
