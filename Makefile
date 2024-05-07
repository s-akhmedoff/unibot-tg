hot_run:
	go run cmd/main.go

build:
	go build -o bin/unibot cmd/main.go

test:
	go test ./...

run:
	./bin/unibot

lint:
	golangci-lint run

cover:
	go test -v -coverprofile=profile.cover ./...