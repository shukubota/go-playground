package main

import (
	"fmt"
	"github.com/shukubota/go-api-template/grpc-example/server/infrastructure"
	ri "github.com/shukubota/go-api-template/grpc-example/server/interfaces"
	"log"
)

func main() {
	mr, err := infrastructure.NewMessageRepository()
	if err != nil {
		log.Fatal(err)
	}

	input := &ri.Message{
		X:    1,
		Y:    2,
		From: "red",
	}
	err = mr.Put(input)
	if err != nil {
		fmt.Println(err)
		fmt.Println("sqs put error")
	}

	messages, err := mr.Get()
	if err != nil {
		fmt.Println(err)
		fmt.Println("sqs get error")
	}

	for _, m := range messages {
		fmt.Println(m.ID)
		fmt.Println(m.From)
		fmt.Println(m.X)
		fmt.Println(m.Y)
	}

	err = mr.Delete(messages)
	if err != nil {
		fmt.Println(err)
		fmt.Println("sqs delete error")
	}

	fmt.Println(messages)
	fmt.Println("-------------message")
}
