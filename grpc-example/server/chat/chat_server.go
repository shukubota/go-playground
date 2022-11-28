package chat_server

import (
	"context"
	"fmt"
	pb "github.com/shukubota/go-api-template/grpc-example/protobuf/server/protobuf"
	"time"
)

type chatServer struct {
	pb.UnimplementedChatServer
	// userId, streamのmap
	clientIdList map[string]pb.Chat_ConnectServer
	data         []*message
}

type message struct {
	from string
	body string
}

func NewChatServer() *chatServer {
	return &chatServer{
		clientIdList: make(map[string]pb.Chat_ConnectServer),
		data:         make([]*message, 0),
	}
}

func (cs *chatServer) Connect(req *pb.ChatConnectRequest, server pb.Chat_ConnectServer) error {
	fmt.Println("----------connect")
	fmt.Println(req)
	fmt.Println(server)
	if cs.clientIdList[req.GetToken()] == nil {
		cs.clientIdList[req.GetToken()] = server
	}
	fmt.Println(cs.clientIdList)
	//count := len(cs.data)
	for {
		// tokenがなければ終了
		if cs.clientIdList[req.GetToken()] == nil {
			return nil
		}
		fmt.Println(len(cs.data))
		//if count < len(cs.data) {
		//	for i := count; i < len(cs.data); i++ {
		//		target := cs.data[i]
		//		err := server.Send(&pb.ChatConnectResponse{
		//			Status: fmt.Sprintf("%v %v", target.body, target.from),
		//		})
		//		fmt.Println(err)
		//		if err != nil {
		//			fmt.Println(err)
		//			return err
		//		}
		//		count++
		//	}
		//}
		time.Sleep(time.Second * 3)
	}
	return nil
}

func (cs *chatServer) SendData(ctx context.Context, req *pb.ChatSendDataRequest) (*pb.ChatSendDataResponse, error) {
	fmt.Println("==============aaaa")
	fmt.Println(req)
	cs.data = append(cs.data, &message{
		from: req.GetFrom(),
		body: req.GetData(),
	})
	for key, stream := range cs.clientIdList {
		fmt.Println(key)
		fmt.Println(stream)
		fmt.Println("============stream")
		if stream == nil {
			continue
		}
		err := stream.Send(&pb.ChatConnectResponse{
			Status: fmt.Sprintf("from %v stream response", key),
		})
		if err != nil {
			fmt.Println(err)
			fmt.Println("send error in sendData")
		}
	}
	return &pb.ChatSendDataResponse{
		Status: "OK",
	}, nil
}
