# qiita-stocker-lambda
[Mindexer](https://www.mindexer.net/) の運用に必要なLambda関数を管理する

## Requirements

* AWS CLI already configured with Administrator permission
* SAM CLI, version 0.39.0
* [Docker installed](https://www.docker.com/community-edition)
* [Golang](https://golang.org)

## Setup process

### Building the project

Goアプリケーションのビルドを行います。

```bash
make build
```

### Packaging and deployment

デプロイ用のS3バケットの作成、ソースコートのZip化、S3バケットへのアップロード、Lambda
のデプロイを行います。

`samconfig.toml` に設定された値を元にデプロイされます。

```bash
make deploy
```
