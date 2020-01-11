package infrastructure

import (
	"github.com/nekochans/qiita-stocker-lambda/stop-fargate/infrastructure/client"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecs"
)

type ECS struct {
	Cluster string
	Service string
}

type ECSRepository interface {
	StopFargate(e ECS) error
}

type ecsRepository struct {
	client *client.ECS
}

func NewECSRepository(c *client.ECS) ECSRepository {
	return &ecsRepository{
		client: c,
	}
}

func (e *ecsRepository) StopFargate(target ECS) error {
	input := ecs.UpdateServiceInput{
		Cluster:      aws.String(target.Cluster),
		Service:      aws.String(target.Service),
		DesiredCount: aws.Int64(0),
	}

	_, err := e.client.ECSAPI.UpdateService(&input)
	if err != nil {
		return err
	}
	return nil
}
