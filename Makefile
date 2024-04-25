.PHONY: build seed run-local clean deploy

build:
	go mod tidy
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bootstrap main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o ./bin/simple-api-go main.go

run-local: build
	env RUNNING_MODE=local ./bin/simple-api-go

seed:
	docker compose up -d
	./schema/schema-seed-data.sh

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	serverless deploy --verbose