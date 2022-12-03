package infrastructure

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type adaptor struct {
	client *dynamodb.DynamoDB
}

type PutData map[string]*dynamodb.AttributeValue
type GetParams map[string]*dynamodb.AttributeValue
type DeleteParams map[string]*dynamodb.AttributeValue
type GetData *dynamodb.GetItemOutput
type QueryData *dynamodb.QueryOutput
type ScanData *dynamodb.ScanOutput

func NewDynamoAdaptor() (*adaptor, error) {
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
	input := &dynamodb.GetItemInput{
		TableName: aws.String(table),
		Key:       p,
	}
	res, err := d.client.GetItem(input)
	fmt.Println(res.Item)
	fmt.Println(err)
	fmt.Println("==============wwww")
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (d *adaptor) Query(table string) (QueryData, error) {
	input := &dynamodb.QueryInput{
		TableName: aws.String(table),
	}
	res, err := d.client.Query(input)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (d *adaptor) Scan(table string) (ScanData, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(table),
	}
	res, err := d.client.Scan(input)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (d *adaptor) Delete(table string, p DeleteParams) error {
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(table),
		Key:       p,
	}
	_, err := d.client.DeleteItem(input)
	if err != nil {
		return err
	}
	return nil
}
