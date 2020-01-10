package app

import (
	"errors"
	"fmt"
	"log"

	"github.com/nekochans/qiita-stocker-lambda/stop-fargate/infrastructure"
)

func StopFargate(rep infrastructure.ECSRepository) error {
	errorCount := 0
	for _, e := range targetECSs() {
		if err := rep.StopFargate(e); err != nil {
			log.Printf("failure to update service : %v", err)
			log.Printf("cluster: %s, service: %s", e.Cluster, e.Service)
			errorCount++
		}
	}

	if errorCount != 0 {
		return errors.New(fmt.Sprintf("failure to stop fargate"))
	}
	return nil
}

func targetECSs() []infrastructure.ECS {
	return []infrastructure.ECS{
		{Cluster: "stg-fargate-api", Service: "stg-fargate-api"},
	}
}
