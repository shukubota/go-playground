package service

import (
	"context"
	"fmt"
	pb "github.com/shukubota/go-api-template/grpc-example/protobuf/server/protobuf"
	"github.com/shukubota/go-api-template/grpc-example/server/infrastructure"
	ri "github.com/shukubota/go-api-template/grpc-example/server/interfaces"
	"time"
)

type chatServer struct {
	pb.UnimplementedChatServer
	cr ri.ConnectionRepository
	// userId, streamのmap
	clientIdList map[string]pb.Chat_ConnectServer
	data         []*message
}

type message struct {
	from string
	x    uint64
	y    uint64
}

func NewChatServer() (*chatServer, error) {
	cr, err := infrastructure.NewConnectionRepository()
	if err != nil {
		return nil, err
	}
	return &chatServer{
		cr:           cr,
		clientIdList: make(map[string]pb.Chat_ConnectServer),
		data:         make([]*message, 0),
	}, nil
}

// 接続stream ここでデータ変更を検知してclientに送信

func (cs *chatServer) Connect(req *pb.ChatConnectRequest, server pb.Chat_ConnectServer) error {
	fmt.Println("----------connect")
	fmt.Println(req)
	fmt.Println(server)
	cs.clientIdList[req.GetToken()] = server
	count := len(cs.data)
	// connectionをdynamoへ
	cs.cr.Put(&ri.Connection{
		ID: req.GetToken(),
	})

	res, err := cs.cr.Get(req.GetToken())
	fmt.Println(err)
	fmt.Println(res)
	fmt.Println("=======================err")

	// ここはgoroutine出して良さげ
	for {
		// tokenがなければ終了
		if cs.clientIdList[req.GetToken()] == nil {
			return nil
		}
		//fmt.Println(len(cs.data))
		if count < len(cs.data) {
			for i := count; i < len(cs.data); i++ {
				target := cs.data[i]
				for un, stream := range cs.clientIdList {
					// 送り主には送らない
					if un == target.from {
						//continue
					}

					err := stream.Send(&pb.ChatConnectResponse{
						From: target.from,
						Data: &pb.DotData{
							X: target.x,
							Y: target.y,
						},
					})
					if err != nil {
						fmt.Println(err)
						return err
					}
				}
				count++
			}
		}
		time.Sleep(time.Millisecond * 200)
	}
	return nil
}

// clientから受け取る

func (cs *chatServer) SendData(ctx context.Context, req *pb.ChatSendDataRequest) (*pb.ChatSendDataResponse, error) {
	data := req.GetData()
	x := data.GetX()
	y := data.GetY()
	cs.data = append(cs.data, &message{
		from: req.GetFrom(),
		x:    x,
		y:    y,
	})
	return &pb.ChatSendDataResponse{
		Status: "OK",
	}, nil
}
