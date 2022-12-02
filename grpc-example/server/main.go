package main

import (
	"flag"
	pb "github.com/shukubota/go-api-template/grpc-example/protobuf/server/protobuf"
	"github.com/shukubota/go-api-template/grpc-example/server/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

var (
	port = flag.Int("port", 50051, "the server port")
)

//type myServer struct {
//	pb.UnimplementedGreeterServer
//	connectedClient any
//}
//
//func (s *myServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
//	fmt.Println("sayhello")
//	return &pb.HelloReply{
//		Message: fmt.Sprintf("hello, %s!", req.GetName()),
//	}, nil
//}
//
//// bidirection stream用
//func (s *myServer) SayHelloBiDirectionalStream(stream pb.Greeter_SayHelloBiDirectionalStreamServer) error {
//	fmt.Println(stream)
//	for {
//		req, err := stream.Recv()
//		fmt.Println(req)
//		if err != nil {
//			if errors.Is(err, io.EOF) {
//				fmt.Println("EOF")
//				return nil
//			}
//			fmt.Println(err)
//			//time.Sleep(time.Second * 3)
//			return err
//		}
//
//		time.Sleep(time.Second * 2)
//		err = stream.Send(&pb.HelloReply{
//			Message: fmt.Sprintf("%v resだよ", req.Name),
//		})
//		if err != nil {
//			fmt.Println("stream send error")
//			fmt.Println(err)
//		}
//	}
//}
//
//// server stream
//func (s *myServer) SayHelloServerStream(req *pb.HelloRequest, stream pb.Greeter_SayHelloServerStreamServer) error {
//	//resCount := 5
//	//for i := 0; i < resCount; i++ {
//	fmt.Println("request from client")
//	i := 0
//	for {
//		i = i + 1
//		if err := stream.Send(&pb.HelloReply{
//			Message: fmt.Sprintf("[%d] Hello, %s!", i, req.GetName()),
//		}); err != nil {
//			return err
//		}
//		if i > 3 {
//			break
//		}
//		time.Sleep(time.Second * 1)
//	}
//	return nil
//}

//func NewMyServer() *myServer {
//	return &myServer{}
//}

func main() {
	port := 50051
	listener, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()

	//hellopb.RegisterGreeterServer(s, NewMyServer())
	cs, err := service.NewDrawingSharingServer()
	if err != nil {
		log.Fatal(err)
	}
	pb.RegisterDrawingShareServer(s, cs)
	reflection.Register(s)

	log.Printf("start gRPC server port: %v", port)
	s.Serve(listener)
}
