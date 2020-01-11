package client

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/aws/aws-sdk-go/service/ecs/ecsiface"
)

type ECS struct {
	ecsiface.ECSAPI
}

func NewECSClient() (*ECS, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1")},
	)
	if err != nil {
		return &ECS{}, err
	}

	return &ECS{
		ECSAPI: ecs.New(sess),
	}, nil
}
