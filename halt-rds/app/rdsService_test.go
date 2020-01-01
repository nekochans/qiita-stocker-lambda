package app

import (
	"testing"

	"github.com/nekochans/qiita-stocker-lambda/halt-rds/infrastructure"
	"github.com/nekochans/qiita-stocker-lambda/halt-rds/infrastructure/client"

	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/aws/aws-sdk-go/service/rds/rdsiface"
)

type mockRDSClient struct {
	rdsiface.RDSAPI
}

func (m *mockRDSClient) StopDBCluster(*rds.StopDBClusterInput) (*rds.StopDBClusterOutput, error) {
	output := &rds.StopDBClusterOutput{}

	return output, nil
}

func TestHaltRDS(t *testing.T) {

	t.Run("success halt RDS clusters", func(t *testing.T) {
		c := &client.RDS{RDSAPI: &mockRDSClient{}}
		rdsRep := infrastructure.NewRDSRepository(c)

		HaltRDS(rdsRep)
	})
}
