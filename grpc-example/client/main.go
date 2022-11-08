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

	//scanner := bufio.NewScanner(os.Stdin)
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

	fmt.Println(conn)
	fmt.Println("----------connection")

	client := hellopb.NewGreeterClient(conn)
	req := &hellopb.HelloRequest{Name: "hoge"}

	res, err := client.SayHello(context.Background(), req)
	fmt.Println(res.Message)
	fmt.Println(res.Message == "hello, hoge!")
}
