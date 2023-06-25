package service

import (
	"context"
	"fmt"
	pb "github.com/shukubota/go-api-template/grpc-example/protobuf/server/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"log"
	"sync"
	"time"
)

type drawingSharingServer struct {
	pb.UnimplementedDrawingShareServer
	//cr ri.ConnectionRepository
	//mr ri.MessageRepository
	// userId, streamのmap
	sci *safeConnectionInfo
}

type greeterServer struct {
	pb.UnimplementedGreeterServer
}

type safeConnectionInfo struct {
	mu sync.Mutex
	v  map[string]pb.DrawingShare_ConnectServer
}

func (sc *safeConnectionInfo) connect(uid string, conn pb.DrawingShare_ConnectServer) {
	sc.mu.Lock()
	sc.v[uid] = conn
	sc.mu.Unlock()
}

func (sc *safeConnectionInfo) disconnect(uid string) {
	sc.mu.Lock()
	delete(sc.v, uid)
	sc.mu.Unlock()
}

func (sc *safeConnectionInfo) value(uid string) pb.DrawingShare_ConnectServer {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	return sc.v[uid]
}

func (sc *safeConnectionInfo) allConnections() map[string]pb.DrawingShare_ConnectServer {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	return sc.v
}

func NewDrawingSharingServer() (*drawingSharingServer, error) {
	//cr, err := infrastructure.NewConnectionRepository()
	//if err != nil {
	//	return nil, err
	//}
	//mr, err := infrastructure.NewMessageRepository()
	//if err != nil {
	//	return nil, err
	//}
	return &drawingSharingServer{
		//cr: cr,
		//mr: mr,
		sci: &safeConnectionInfo{
			v: make(map[string]pb.DrawingShare_ConnectServer),
		},
	}, nil
}

// 接続stream ここでデータ変更を検知してclientに送信

func (cs *drawingSharingServer) Connect(req *pb.ConnectRequest, server pb.DrawingShare_ConnectServer) error {
	u := req.GetUser()
	cs.sci.connect(u, server)
	defer cs.sci.disconnect(u)

	for {
		if cs.sci.value(u) == nil {
			return nil
		}
	}
	return nil
}

func (cs *drawingSharingServer) DisConnect(ctx context.Context, req *pb.DisConnectRequest) (*pb.DisConnectResponse, error) {
	u := req.GetUser()
	cs.sci.disconnect(u)
	return &pb.DisConnectResponse{
		Status: "OK",
	}, nil
}

// clientから受け取る

func (cs *drawingSharingServer) SendDrawing(ctx context.Context, req *pb.SendDrawingRequest) (*pb.SendDrawingResponse, error) {
	data := req.GetData()
	x := data.GetX()
	y := data.GetY()
	from := req.GetFrom()
	for _, stream := range cs.sci.allConnections() {
		err := stream.Send(&pb.ConnectResponse{
			From: from,
			Data: &pb.DotData{
				X: x,
				Y: y,
			},
		})
		if err != nil {
			log.Printf("%+v", err)
			continue
		}
	}

	return &pb.SendDrawingResponse{
		Status: "OK",
	}, nil
}

func NewGreeterServer() (*greeterServer, error) {
	return &greeterServer{}, nil
}

func (hs *greeterServer) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println(request.GetName())
	fmt.Println("=======================")

	return &pb.HelloReply{
		Message: "hoge",
	}, nil
	//return nil, errors.New("aaa")
	address := "127.0.0.1:50051"

	conn, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                2 * time.Second, // 活動がなくなってから PING を送るまでの間隔
			Timeout:             2 * time.Second, // PING 応答を待つ時間
			PermitWithoutStream: true,            // アクティブなストリームがないときも probe を送るかどうか
		}),
	)
	if err != nil {
		log.Fatal("connection fail")
		fmt.Println(err)
		return nil, err
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	r, err := client.SayHello(ctx, &pb.HelloRequest{
		Name: "test",
	})
	if err != nil {
		log.Printf("%+v\n", err)
		return nil, err
	}

	fmt.Println(r)

	return &pb.HelloReply{
		Message: "hoge",
	}, nil
}

func (hs *greeterServer) SayHelloBiDirectionalStream(server pb.Greeter_SayHelloBiDirectionalStreamServer) error {
	return nil
}

func (hs *greeterServer) SayHelloServerStream(request *pb.HelloRequest, server pb.Greeter_SayHelloServerStreamServer) error {
	return nil
}
