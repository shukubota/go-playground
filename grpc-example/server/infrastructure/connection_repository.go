package infrastructure

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	ri "github.com/shukubota/go-api-template/grpc-example/server/interfaces"
)

type connectionRepository struct {
	db *adaptor
}

func NewConnectionRepository() (ri.ConnectionRepository, error) {
	a, err := NewDynamoAdaptor()
	if err != nil {
		return nil, err
	}
	return &connectionRepository{
		db: a,
	}, nil
}

func (cr *connectionRepository) Get(ui string) (*ri.Connection, error) {
	input := GetParams{
		"user": {
			S: aws.String(ui),
		},
	}
	r, err := cr.db.Get("connections", input)
	if err != nil {
		return nil, err
	}
	u := r.Item["user"]
	if u == nil {
		return nil, nil
	}
	return &ri.Connection{
		ID: *r.Item["user"].S,
	}, nil
}

func (cr *connectionRepository) Put(c *ri.Connection) error {
	fmt.Println(c)
	input := PutData{
		"user": {
			S: aws.String(c.ID),
		},
	}
	err := cr.db.Put("connections", input)
	if err != nil {
		return err
	}
	return nil
}

func (cr *connectionRepository) Delete(c *ri.Connection) error {
	fmt.Println(c)
	input := DeleteParams{
		"user": {
			S: aws.String(c.ID),
		},
	}
	err := cr.db.Delete("connections", input)
	if err != nil {
		return err
	}
	return nil
}
