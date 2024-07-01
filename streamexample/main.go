package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
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

	// io.Reader満たしてない
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
			return nil
		}

		if err != nil {
			return err
		}

		m := response.Choices[0].Delta.Content

		err = server.Send(&healthpb.CheckResponse{
			Message: m,
		})
		if err != nil {
			return err
		}
	}

	//// azure openai
	//azureOpenAIKey := os.Getenv("AOAI_COMPLETIONS_API_KEY")
	//modelDeploymentID := os.Getenv("AOAI_COMPLETIONS_MODEL")
	//
	//azureOpenAIEndpoint := os.Getenv("AOAI_COMPLETIONS_ENDPOINT")
	//
	//if azureOpenAIKey == "" || modelDeploymentID == "" || azureOpenAIEndpoint == "" {
	//	fmt.Fprintf(os.Stderr, "Skipping example, environment variables missing\n")
	//	return errors.New("environment variables missing")
	//}
	//
	//keyCredential := azcore.NewKeyCredential(azureOpenAIKey)
	//
	//c, err := azopenai.NewClientWithKeyCredential(azureOpenAIEndpoint, keyCredential, nil)
	//
	//if err != nil {
	//	// TODO: Update the following line with your application specific error handling logic
	//	log.Printf("ERROR: %s", err)
	//	return err
	//}
	//
	//messages := []azopenai.ChatRequestMessageClassification{
	//	&azopenai.ChatRequestUserMessage{Content: azopenai.NewChatRequestUserMessageContent("Can you help me?")},
	//}
	//
	//resp, err := c.GetChatCompletionsStream(context.TODO(), azopenai.ChatCompletionsOptions{
	//	//Prompt:         []string{"What is Azure OpenAI, in 20 words or less?"},
	//	Messages:       messages,
	//	MaxTokens:      to.Ptr(int32(2048)),
	//	Temperature:    to.Ptr(float32(0.0)),
	//	DeploymentName: &modelDeploymentID,
	//}, nil)
	//
	//if err != nil {
	//	log.Printf("ERROR: %s", err)
	//	return err
	//}
	//
	//defer resp.ChatCompletionsStream.Close()
	//
	//for {
	//	entry, err := resp.ChatCompletionsStream.Read()
	//
	//	if errors.Is(err, io.EOF) {
	//		fmt.Fprintf(os.Stderr, "\n*** No more completions ***\n")
	//		break
	//	}
	//
	//	if err != nil {
	//		//  TODO: Update the following line with your application specific error handling logic
	//		log.Printf("ERROR: %s", err)
	//		return err
	//	}
	//
	//	for _, choice := range entry.Choices {
	//		fmt.Fprintf(os.Stderr, "Result: %+v\n", choice)
	//	}
	//}

	return nil
}
