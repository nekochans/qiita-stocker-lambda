package app

import (
	"github.com/nekochans/qiita-stocker-lambda/halt-rds/infrastructure"
)

func HaltRDS(rep infrastructure.RDSRepository) {
	for _, ID := range targetRDSClusterIDs() {
		rep.HaltRDSClusters(ID)
	}
}

func targetRDSClusterIDs() []string {
	return []string{"stg-database-cluster"}
}
