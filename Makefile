REGION    := ap-northeast-1
PROFILE   := qiita-stocker-dev
STACK_NAME := dev-qiita-stocker-lambda

.PHONY: deps clean build haltrds

clean:
	rm -rf ./bin/halt-rds

build: clean haltrds stopfg

test:
	go test -v ./...

haltrds:
	GOOS=linux GOARCH=amd64 go build -o bin/halt-rds ./halt-rds
	chmod +x bin/halt-rds

stopfg:
	GOOS=linux GOARCH=amd64 go build -o bin/stop-fargate ./stop-fargate
	chmod +x bin/stop-fargate

validate:
	sam validate \
		--profile $(PROFILE) \
		--region $(REGION)

deploy:
	sam deploy \
		--profile $(PROFILE)

delete-lambda-stack:
	aws cloudformation delete-stack \
		--stack-name $(STACK_NAME) \
		--profile $(PROFILE) \
		--region $(REGION)
