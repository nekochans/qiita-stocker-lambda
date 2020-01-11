package app

import (
	"errors"
	"testing"

	"github.com/nekochans/qiita-stocker-lambda/stop-fargate/infrastructure"
	"github.com/nekochans/qiita-stocker-lambda/stop-fargate/infrastructure/client"

	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/aws/aws-sdk-go/service/ecs/ecsiface"
)

type mockECSClient struct {
	ecsiface.ECSAPI
}

type mockErrorECSClient struct {
	ecsiface.ECSAPI
}

func (m *mockECSClient) UpdateService(*ecs.UpdateServiceInput) (*ecs.UpdateServiceOutput, error) {
	output := &ecs.UpdateServiceOutput{}

	return output, nil
}

func (m *mockErrorECSClient) UpdateService(*ecs.UpdateServiceInput) (*ecs.UpdateServiceOutput, error) {
	output := &ecs.UpdateServiceOutput{}

	return output, errors.New("")
}

func TestStopFargate(t *testing.T) {
	t.Run("success stop fargate", func(t *testing.T) {
		c := &client.ECS{ECSAPI: &mockECSClient{}}
		rdsRep := infrastructure.NewECSRepository(c)

		err := StopFargate(rdsRep)
		if err != nil {
			t.Errorf("error should not be returned")
		}
	})

	t.Run("failure stop fargate", func(t *testing.T) {
		c := &client.ECS{ECSAPI: &mockErrorECSClient{}}
		rdsRep := infrastructure.NewECSRepository(c)

		err := StopFargate(rdsRep)
		if err == nil {
			t.Errorf("error should not be returned")
		}
	})

}
