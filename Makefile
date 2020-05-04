.PHONY: build clean deploy

build:
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/get-all-pairs aws/get-all-pairs/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/get-pairs aws/get-pairs/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/get-meeting-structure aws/get-meeting-structure/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/get-meetings aws/get-meetings/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy-mentornity: clean build
	serverless deploy --stage mentornity --region eu-central-1 --verbose

teardown-mentornity: clean
	serverless remove --stage mentornity --region eu-central-1 --verbose
