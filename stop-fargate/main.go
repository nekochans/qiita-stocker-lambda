package main

import (
	"github.com/nekochans/qiita-stocker-lambda/stop-fargate/app"
	"github.com/nekochans/qiita-stocker-lambda/stop-fargate/infrastructure"
	"github.com/nekochans/qiita-stocker-lambda/stop-fargate/infrastructure/client"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	rep infrastructure.ECSRepository
)

func init() {
	client, err := client.NewECSClient()
	if err != nil {
		panic("failed initialize client: " + err.Error())
	}
	rep = infrastructure.NewECSRepository(client)
}

func handler(events.CloudWatchEvent) error {
	if err := app.StopFargate(rep); err != nil {
		return err
	}
	return nil
}

func main() {
	lambda.Start(handler)
}
