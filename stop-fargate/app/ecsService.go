package app

import (
	"github.com/nekochans/qiita-stocker-lambda/stop-fargate/infrastructure"
)

func StopFargate(rep infrastructure.ECSRepository) {
	for _, e := range targetECSs() {
		rep.StopFargate(e)
	}
}

func targetECSs() []infrastructure.ECS {
	return []infrastructure.ECS{
		{Cluster: "stg-fargate-api", Service: "stg-fargate-api"},
	}
}
