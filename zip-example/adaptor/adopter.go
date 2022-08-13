package adaptor

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
)

type Adapter struct {
	s3 *s3.S3
}

func NewAdapter() (*Adapter, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String("us-east-1"),
		Endpoint: aws.String("http://localhost:4566"),
	})
	if err != nil {
		return nil, err
	}
	return &Adapter{
		s3: s3.New(sess),
	}, nil
}

func (s *Adapter) Upload(body io.ReadSeeker) error {
	i := &s3.PutObjectInput{
		Bucket: aws.String("test"),
		Key:    aws.String("test/fugaa.zip"),
		Body:   body,
	}

	_, err := s.s3.PutObject(i)

	if err != nil {
		return err
	}
	return nil
}
