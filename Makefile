all:
	go run ./cmd/webapi/main.go

test:
	go run ./cmd/webapi/main.go

test-cov:
	go run ./cmd/webapi/main.go

unit-test:
	go run ./cmd/webapi/main.go

integration-test:
	go run ./cmd/webapi/main.go

lint:
	go run ./cmd/webapi/main.go
