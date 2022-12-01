package infrastructure

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	ri "github.com/shukubota/go-api-template/grpc-example/server/interfaces"
	"strconv"
)

type messageRepositoryImpl struct {
	db *sqsAdaptor
}

func NewMessageRepository() (ri.MessageRepository, error) {
	a, err := NewSQSAdaptor()
	if err != nil {
		return nil, err
	}
	return &messageRepositoryImpl{
		db: a,
	}, nil
}

func (cr *messageRepositoryImpl) Put(m *ri.Message) error {
	input := sqsSendData{
		"from": &sqs.MessageAttributeValue{
			StringValue: aws.String(m.From),
			DataType:    aws.String("String"),
		},
		"x": &sqs.MessageAttributeValue{
			StringValue: aws.String(fmt.Sprintf("%v", m.X)),
			DataType:    aws.String("Number"),
		},
		"y": &sqs.MessageAttributeValue{
			StringValue: aws.String(fmt.Sprintf("%v", m.Y)),
			DataType:    aws.String("Number"),
		},
	}
	err := cr.db.Put("messages", input)
	if err != nil {
		return err
	}
	return nil
}

// キューから1個取り出す

func (cr *messageRepositoryImpl) Get() ([]*ri.Message, error) {
	m, err := cr.db.Get("messages")
	if err != nil {
		return nil, err
	}
	fmt.Println(len(m.Messages))

	messages := make([]*ri.Message, 0)
	for _, message := range m.Messages {
		x, err := strconv.ParseInt(*message.MessageAttributes["x"].StringValue, 10, 64)
		if err != nil {
			return messages, err
		}
		y, err := strconv.ParseInt(*message.MessageAttributes["y"].StringValue, 10, 64)
		if err != nil {
			return messages, err
		}
		messages = append(messages, &ri.Message{
			ID:   *message.MessageId,
			Rh:   *message.ReceiptHandle,
			From: *message.MessageAttributes["from"].StringValue,
			X:    int(x),
			Y:    int(y),
		})
	}

	return messages, nil
}

func (cr *messageRepositoryImpl) Delete(messages []*ri.Message) error {
	var deleteData []sqsDeleteData
	for _, m := range messages {
		e := sqsDeleteData{
			Id:            aws.String(m.ID),
			ReceiptHandle: aws.String(m.Rh),
		}
		deleteData = append(deleteData, e)
	}
	err := cr.db.Delete("messages", deleteData)
	if err != nil {
		return err
	}
	return nil
}
