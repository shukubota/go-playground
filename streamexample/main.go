package main

import (
	"context"
	"fmt"
	healthpb "github.com/shukubota/go-playground/streamexample/gen/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
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
