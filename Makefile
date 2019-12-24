.PHONY: deps clean build haltrds

clean:
	rm -rf ./bin/halt-rds

build: clean haltrds

haltrds:
	GOOS=linux GOARCH=amd64 go build -o bin/halt-rds ./halt-rds
	chmod +x bin/halt-rds
