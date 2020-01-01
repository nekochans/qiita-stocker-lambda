package main

import (
	"github.com/nekochans/qiita-stocker-lambda/halt-rds/app"
	"github.com/nekochans/qiita-stocker-lambda/halt-rds/infrastructure"
	"github.com/nekochans/qiita-stocker-lambda/halt-rds/infrastructure/client"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	rdsRep infrastructure.RDSRepository
)

func init() {
	client, err := client.NewRDSClient()
	if err != nil {
		panic("failed initialize client: " + err.Error())
	}
	rdsRep = infrastructure.NewRDSRepository(client)
}

func handler(request events.CloudWatchEvent) error {
	app.HaltRDS(rdsRep)
	return nil
}

func main() {
	lambda.Start(handler)
}
