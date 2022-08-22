package main

import (
	"context"
	hellopb "example/hello/grpc-example/protobuf"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

var (
	port = flag.Int("port", 50051, "the server port")
)

type GreeterServer interface {
	SayHello(context.Context, *hellopb.HelloRequest) (*hellopb.HelloReply, error)
	mustEmbedUnimplementedGreeterServer()
}

type myServer struct {
	hellopb.UnimplementedGreeterServer
}

func (s *myServer) SayHello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloReply, error) {
	return &hellopb.HelloReply{
		Message: fmt.Sprintf("hello, %s!", req.GetName()),
	}, nil
}

func NewMyServer() *myServer {
	return &myServer{}
}

func main() {
	fmt.Println("grpc-example")
	port := 50051
	listner, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	hellopb.RegisterGreeterServer(s, NewMyServer())
	go func() {
		log.Printf("start gRPC server port: %v", port)
		s.Serve(listner)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}
