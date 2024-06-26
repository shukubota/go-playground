package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/sashabaranov/go-openai"
	healthpb "github.com/shukubota/go-playground/streamexample/gen/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	err := Main()
	if err != nil {
		log.Fatal(err)
	}
}

func Main() error {
	port := ":50051"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	healthpb.RegisterHealthServiceServer(grpcServer, &healthServer{})
	reflection.Register(grpcServer)

	log.Printf("start gRPC server port: %v", port)
	grpcServer.Serve(listener)

	return nil
}

type healthServer struct{}

func (s *healthServer) Check(ctx context.Context, req *healthpb.CheckRequest) (*healthpb.CheckResponse, error) {
	fmt.Println(req.Message)
	return &healthpb.CheckResponse{
		Message: "OK",
	}, nil
}

func (s *healthServer) CheckStream(request *healthpb.CheckRequest, server healthpb.HealthService_CheckStreamServer) error {
	client := openai.NewClient(os.Getenv("OPEN_AI_API_KEY"))
	ctx := context.Background()
	stream, err := client.CreateChatCompletionStream(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Hello!",
				},
			},
			Stream: true,
		},
	)
	if err != nil {
		return err
	}

	defer stream.Close()

	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("\nStream finished")
			return nil
		}

		if err != nil {
			fmt.Printf("\nStream error: %v\n", err)
			return err
		}

		m := response.Choices[0].Delta.Content

		err = server.Send(&healthpb.CheckResponse{
			Message: m,
		})
		if err != nil {
			return err
		}

		fmt.Printf(response.Choices[0].Delta.Content)
	}

	return nil
}
