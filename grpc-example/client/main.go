package main

import (
	"context"
	"fmt"
	hellopb "github.com/shukubota/go-api-template/grpc-example/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	err = requestUnary(conn)
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
