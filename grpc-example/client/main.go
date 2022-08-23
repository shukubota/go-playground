package main

import (
	"bufio"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

func main() {
	fmt.Println("start grpc client")

	scanner := bufio.NewScanner(os.Stdin)
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

}
