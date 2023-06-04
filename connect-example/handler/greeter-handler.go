package handler

import (
	"connect-example/gen/greet/v1"
	"context"
	"github.com/bufbuild/connect-go"
	"log"
)

type greetHandler struct{}

type GreetHandlerIF interface {
	Greet(ctx context.Context, req *connect.Request[greetv1.GreetRequest]) (*connect.Response[greetv1.GreetResponse], error)
}

func NewGreetHandler() GreetHandlerIF {
	return &greetHandler{}
}

func (s *greetHandler) Greet(ctx context.Context, req *connect.Request[greetv1.GreetRequest]) (*connect.Response[greetv1.GreetResponse], error) {
	log.Println("RequestHeaders: ", req.Header())

	res := connect.NewResponse(&greetv1.GreetResponse{
		Greeting: "Hello " + req.Msg.Name,
	})
	return res, nil
}
