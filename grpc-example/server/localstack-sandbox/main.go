package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/google/uuid"
	"log"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String("ap-northeast-1"),
		Endpoint: aws.String("http://localhost:4566"),
	})
	db := dynamodb.New(sess)

	//di := &dynamodb.DeleteItemInput{
	//	TableName: aws.String("connections"),
	//	Key: map[string]*dynamodb.AttributeValue{
	//		"user": {
	//			S: aws.String("green"),
	//		},
	//	},
	//}
	//
	//_, err = db.DeleteItem(di)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//// connection管理 connections
	//param := &dynamodb.PutItemInput{
	//	TableName: aws.String("connections"),
	//	Item: map[string]*dynamodb.AttributeValue{
	//		"user": {
	//			S: aws.String("green"), //データ型(String:S)
	//		},
	//	},
	//}
	//_, err = db.PutItem(param)
	//
	//if err != nil {
	//	fmt.Println(err.Error())
	//}

	uuid1 := uuid.New()
	uuid2 := uuid.New()
	fmt.Println(uuid1)
	fmt.Println(uuid2)

	//di := &dynamodb.DeleteItemInput{
	//	TableName: aws.String("messages"),
	//	Key: map[string]*dynamodb.AttributeValue{
	//		"user": {
	//			S: aws.String("green"),
	//		},
	//	},
	//}
	//
	//_, err = db.DeleteItem(di)
	//if err != nil {
	//	log.Fatal(err)
	//}

	// message管理 messages
	mi := &dynamodb.PutItemInput{
		TableName: aws.String("messages"),
		Item: map[string]*dynamodb.AttributeValue{
			"uuid": {
				S: aws.String(uuid1.String()),
			},
			"user": {
				S: aws.String("green"),
			},
			"x": {
				N: aws.String("3"),
			},
			"y": {
				N: aws.String("5"),
			},
		},
	}
	_, err = db.PutItem(mi)

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("putitem error")
	}
	mi = &dynamodb.PutItemInput{
		TableName: aws.String("messages"),
		Item: map[string]*dynamodb.AttributeValue{
			"uuid": {
				S: aws.String(uuid2.String()),
			},
			"user": {
				S: aws.String("red"),
			},
			"x": {
				N: aws.String("4"),
			},
			"y": {
				N: aws.String("8"),
			},
		},
	}
	_, err = db.PutItem(mi)

	if err != nil {
		fmt.Println(err.Error())
	}

	//query := &dynamodb.GetItemInput{
	//	TableName: aws.String("connections"),
	//
	//	Key: map[string]*dynamodb.AttributeValue{
	//		"user": {
	//			S: aws.String("green"),
	//		},
	//	},
	//}
	//res, err := db.GetItem(query)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(res)
	//fmt.Println(*res.Item["user"].S)

	//mq := &dynamodb.GetItemInput{
	//	TableName: aws.String("messages"),
	//
	//	Key: map[string]*dynamodb.AttributeValue{
	//		"user": {
	//			S: aws.String("green"),
	//		},
	//	},
	//}
	//res, err = db.GetItem(mq)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}

	//fmt.Println(res)
	//fmt.Println(*res.Item["x"].N)
	//fmt.Println(*res.Item["y"].N)

	ms := &dynamodb.ScanInput{
		TableName: aws.String("messages"),
	}
	scanRes, err := db.Scan(ms)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(scanRes.Items)
	for _, v := range scanRes.Items {
		fmt.Println(*v["uuid"].S)
		fmt.Println(*v["x"].N)
		fmt.Println(*v["y"].N)
		fmt.Println(*v["user"].S)
	}
	fmt.Println("===================scanres")

	fmt.Println("============ddd")
}
