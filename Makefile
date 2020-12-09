all:
	go run ./cmd/webapi/main.go

test:
	go test ./...

lint:
	golint ./...
