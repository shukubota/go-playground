package main

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	pb "github.com/shukubota/go-api-template/grpc-example/protobuf/server/protobuf"
	"github.com/shukubota/go-api-template/grpc-example/server/interceptor"
	"github.com/shukubota/go-api-template/grpc-example/server/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

//var (
//	port = flag.Int("port", 50051, "the server port")
//)

func main() {
	port := 50051
	listener, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatal(err)
	}

	interceptor := grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(interceptor.UnaryInterceptor))
	s := grpc.NewServer(interceptor)

	cs, err := service.NewDrawingSharingServer()
	if err != nil {
		log.Fatal(err)
	}
	hs, err := service.NewGreeterServer()
	if err != nil {
		log.Fatal(err)
	}

	pb.RegisterDrawingShareServer(s, cs)
	pb.RegisterGreeterServer(s, hs)
	reflection.Register(s)

	log.Printf("start gRPC server port: %v", port)
	s.Serve(listener)
}
