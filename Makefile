hot_run:
	go run cmd/main.go

build:
	go build -o bin/unibot cmd/main.go

test:
	go test ./...

run:
	./bin/unibot

valgrind:
	valgrind --leak-check=full --track-origins=yes -s bin/unibot

bench:
	go test -bench ./..