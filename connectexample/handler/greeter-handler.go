package handler

import (
	v1 "connectexample/gen/greet/v1"
	"connectexample/gen/greet/v1/v1connect"
	"connectrpc.com/connect"
	"context"
	"log"
)

type greetHandler struct{}

func NewGreetHandler() v1connect.GreetServiceHandler {
	return &greetHandler{}
}

func (s *greetHandler) Greet(ctx context.Context, req *connect.Request[v1.GreetRequest]) (*connect.Response[v1.GreetResponse], error) {
	log.Println("RequestHeaders: ", req.Header())

	res := connect.NewResponse(&v1.GreetResponse{
		Greeting: "Hello " + req.Msg.Name,
	})
	return res, nil
}

func (s *greetHandler) Echo(ctx context.Context, req *connect.Request[v1.StringMessage]) (*connect.Response[v1.StringMessage], error) {
	log.Println("RequestHeaders: ", req.Header())

	res := connect.NewResponse(&v1.StringMessage{
		Value: req.Msg.Value,
	})
	return res, nil
}
