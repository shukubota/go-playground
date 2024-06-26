package main

import (
	"context"
	"errors"
	"fmt"
	hellopb "github.com/shukubota/go-playground/grpc-example/protobuf/server/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"time"
)

func main() {
	fmt.Println("start grpc client")
	address := "127.0.0.1:50051"

	// unary 用のconnection
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

	// stream用
	ctx := context.Background()
	client := hellopb.NewGreeterClient(conn)
	stream, err := client.SayHelloBiDirectionalStream(ctx)

	fmt.Println(stream)
	fmt.Println("========stream")

	if err != nil {
		log.Fatalf("%v\n", err)
	}

	// unary
	//err = requestUnary(conn)
	// server stream
	//err = requestServerStream(conn)

	// bidirectional stream
	streamRequestErrorChannel := make(chan error)
	streamReceiveErrorChannel := make(chan error)
	//err = requestBiDirectionalStreaming(stream)
	go requestBiDirectionalStreaming(streamRequestErrorChannel, stream)
	go receiveBiDirectionStreaming(streamReceiveErrorChannel, stream)

	err = <-streamRequestErrorChannel
	receiveError := <-streamReceiveErrorChannel
	fmt.Println(fmt.Sprintf("error %v", err))
	fmt.Println(fmt.Sprintf("reveiveError: %v", receiveError))

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

func requestBiDirectionalStreaming(channel chan<- error, stream hellopb.Greeter_SayHelloBiDirectionalStreamClient) {
	defer close(channel)
	fmt.Println("requestBiDirectionalStreaming")
	for i := 0; i < 3; i++ {
		err := stream.Send(&hellopb.HelloRequest{
			Name: fmt.Sprintf("%v\n", i),
		})
		time.Sleep(time.Second * 1)
		fmt.Println(err)
		if err != nil {
			break
		}
	}
	// ここでcloseしないとlistenし続ける
	err := stream.CloseSend()
	fmt.Println("----closeSend")
	fmt.Println(err)
	if err != nil {
		channel <- err
		//return err
	}
	//return nil
	channel <- nil
	return
}

func receiveBiDirectionStreaming(channel chan<- error, stream hellopb.Greeter_SayHelloBiDirectionalStreamClient) {
	defer close(channel)
	fmt.Println("receiveBiDirectionStreaming")
	for {
		res, err := stream.Recv()
		fmt.Println(fmt.Sprintf("receive result: %v", res))
		fmt.Println(fmt.Sprintf("error %v", err))
		if err != nil {
			channel <- err
			break
		}
	}
	channel <- nil
	return
}
