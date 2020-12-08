all:
	go run ./cmd/webapi/main.go

test:
	go test ./...

test-cov:
	go run ./cmd/webapi/main.go

unit-test:
	go test ./...

integration-test:
	go test ./...

lint:
	go run ./cmd/webapi/main.go
