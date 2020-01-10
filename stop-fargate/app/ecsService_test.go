package app

import (
	"testing"

	"github.com/nekochans/qiita-stocker-lambda/stop-fargate/infrastructure"
	"github.com/nekochans/qiita-stocker-lambda/stop-fargate/infrastructure/client"

	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/aws/aws-sdk-go/service/ecs/ecsiface"
)

type mockECSClient struct {
	ecsiface.ECSAPI
}

func (m *mockECSClient) UpdateService(*ecs.UpdateServiceInput) (*ecs.UpdateServiceOutput, error) {
	output := &ecs.UpdateServiceOutput{}

	return output, nil
}

func TestStopFargate(t *testing.T) {
	t.Run("success stop fargate", func(t *testing.T) {
		c := &client.ECS{ECSAPI: &mockECSClient{}}
		rdsRep := infrastructure.NewECSRepository(c)

		StopFargate(rdsRep)
	})
}
