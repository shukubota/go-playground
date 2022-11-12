package main

import (
	"context"
	"errors"
	"fmt"
	hellopb "github.com/shukubota/go-api-template/grpc-example/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
)

func main() {
	fmt.Println("start grpc client")
	address := "127.0.0.1:50051"

	conn, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("connection fail")
		return
	}
	defer conn.Close()

	//err = requestUnary(conn)
	err = requestServerStream(conn)

	if err != nil {
		log.Fatalf("%v\n", err)
	}
}

func requestUnary(conn *grpc.ClientConn) error {
	client := hellopb.NewGreeterClient(conn)
	req := &hellopb.HelloRequest{Name: "hoge"}

	res, err := client.SayHello(context.Background(), req)
	if err != nil {
		return err
	}
	fmt.Println(res.Message)
	return nil
}

// server Streamでリクエスト
func requestServerStream(conn *grpc.ClientConn) error {
	client := hellopb.NewGreeterClient(conn)
	req := &hellopb.HelloRequest{
		Name: "test",
	}
	stream, err := client.SayHelloServerStream(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return err
	}

	for {
		fmt.Println("---------------999")
		res, err := stream.Recv()
		fmt.Println(res)
		fmt.Println(err)
		if errors.Is(err, io.EOF) {
			fmt.Println("all the responses have already received.")
			break
		}

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res)
	}
	return nil
}
