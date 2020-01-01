package infrastructure

import (
	"log"

	"github.com/nekochans/qiita-stocker-lambda/halt-rds/infrastructure/client"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rds"
)

type RDSRepository interface {
	HaltRDSClusters(targetCluster string)
}

type rdsRepository struct {
	client *client.RDS
}

func NewRDSRepository(c *client.RDS) RDSRepository {
	return &rdsRepository{
		client: c,
	}
}

func (r *rdsRepository) HaltRDSClusters(ID string) {
	input := rds.StopDBClusterInput{
		DBClusterIdentifier: aws.String(ID),
	}

	_, err := r.client.RDSAPI.StopDBCluster(&input)
	if err != nil {
		log.Printf("failer to stop DB Cluster: %v", err)
	}
}
