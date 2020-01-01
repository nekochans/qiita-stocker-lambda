package client

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/aws/aws-sdk-go/service/rds/rdsiface"
)

type RDS struct {
	rdsiface.RDSAPI
}

func NewRDSClient() (*RDS, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1")},
	)
	if err != nil {
		return &RDS{}, err
	}

	return &RDS{
		RDSAPI: rds.New(sess),
	}, nil
}
