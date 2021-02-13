hot_run:
	go run cmd/main.go

build:
	go build -o bin/unibot cmd/main.go

test:
	go test ./...

run:
	./bin/unibot

cover:
	go test -v -coverprofile=profile.cov ./...