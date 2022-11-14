package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	hellopb "github.com/shukubota/go-api-template/grpc-example/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

var (
	port = flag.Int("port", 50051, "the server port")
)

type myServer struct {
	hellopb.UnimplementedGreeterServer
}

func (s *myServer) SayHello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloReply, error) {
	fmt.Println("sayhello")
	return &hellopb.HelloReply{
		Message: fmt.Sprintf("hello, %s!", req.GetName()),
	}, nil
}

// bidirection stream用
func (s *myServer) SayHelloBiDirectionalStream(stream hellopb.Greeter_SayHelloBiDirectionalStreamServer) error {
	fmt.Println(stream)
	for {
		req, err := stream.Recv()
		fmt.Println(req)
		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Println("EOF")
				return nil
			}
			fmt.Println(err)
			//time.Sleep(time.Second * 3)
			return err
		}

		time.Sleep(time.Second * 2)
		err = stream.Send(&hellopb.HelloReply{
			Message: fmt.Sprintf("%v resだよ", req.Name),
		})
		if err != nil {
			fmt.Println("stream send error")
			fmt.Println(err)
		}
	}
}

func (s *myServer) SayHelloServerStream(req *hellopb.HelloRequest, stream hellopb.Greeter_SayHelloServerStreamServer) error {
	//resCount := 5
	//for i := 0; i < resCount; i++ {
	i := 0
	for {
		i = i + 1
		if err := stream.Send(&hellopb.HelloReply{
			Message: fmt.Sprintf("[%d] Hello, %s!", i, req.GetName()),
		}); err != nil {
			return err
		}
		if i > 3 {
			break
		}
		time.Sleep(time.Second * 1)
	}
	return nil
}

func NewMyServer() *myServer {
	return &myServer{}
}

func main() {
	fmt.Println("grpc-example")
	port := 50051
	listner, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	hellopb.RegisterGreeterServer(s, NewMyServer())
	reflection.Register(s)

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
