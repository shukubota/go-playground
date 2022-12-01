package infrastructure

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type sqsAdaptor struct {
	client *sqs.SQS
}

type sqsSendData map[string]*sqs.MessageAttributeValue
type sqsReceiveData *sqs.ReceiveMessageOutput
type sqsDeleteData sqs.DeleteMessageBatchRequestEntry

func NewSQSAdaptor() (*sqsAdaptor, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String("ap-northeast-1"),
		Endpoint: aws.String("http://localhost:4566"),
	})
	if err != nil {
		return nil, err
	}

	cl := sqs.New(sess)
	return &sqsAdaptor{
		client: cl,
	}, nil
}

func (d *sqsAdaptor) Put(queue string, data sqsSendData) error {
	input := &sqs.SendMessageInput{
		QueueUrl:          aws.String(queue),
		MessageAttributes: data,
		MessageBody:       aws.String(""),
	}

	_, err := d.client.SendMessage(input)

	if err != nil {
		return err
	}

	return nil
}

func (d *sqsAdaptor) Get(queue string) (sqsReceiveData, error) {
	input := &sqs.ReceiveMessageInput{
		QueueUrl: aws.String(queue),
	}
	input.SetMessageAttributeNames([]*string{aws.String("All")})
	input.SetMaxNumberOfMessages(10)
	res, err := d.client.ReceiveMessage(input)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (d *sqsAdaptor) Delete(queue string, entries []sqsDeleteData) error {
	var es []*sqs.DeleteMessageBatchRequestEntry
	for _, e := range entries {
		es = append(es, &sqs.DeleteMessageBatchRequestEntry{
			Id:            e.Id,
			ReceiptHandle: e.ReceiptHandle,
		})
	}
	input := &sqs.DeleteMessageBatchInput{
		QueueUrl: aws.String(queue),
		Entries:  es,
	}
	_, err := d.client.DeleteMessageBatch(input)
	if err != nil {
		return err
	}
	return nil
}
