package infrastructure

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type adaptor struct {
	client *dynamodb.DynamoDB
}

type PutData map[string]*dynamodb.AttributeValue
type GetParams map[string]*dynamodb.AttributeValue
type GetData *dynamodb.GetItemOutput

func NewAdaptor() (*adaptor, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String("ap-northeast-1"),
		Endpoint: aws.String("http://localhost:4566"),
	})
	if err != nil {
		return nil, err
	}

	db := dynamodb.New(sess)
	return &adaptor{
		client: db,
	}, nil
}

func (d *adaptor) Put(table string, data PutData) error {
	input := &dynamodb.PutItemInput{
		TableName: aws.String(table),
		Item:      data,
	}
	_, err := d.client.PutItem(input)

	if err != nil {
		return err
	}
	return nil
}

func (d *adaptor) Get(table string, p GetParams) (GetData, error) {
	query := &dynamodb.GetItemInput{
		TableName: aws.String(table),
		Key:       p,
	}
	res, err := d.client.GetItem(query)
	if err != nil {
		return nil, err
	}
	return res, nil
}
