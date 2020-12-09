FROM golang:1.13.7

WORKDIR /go/src/cookify

COPY . .

RUN go get -d -v ./...

RUN go build -o cookify ./cmd/webapi/main.go

ENTRYPOINT ./cookify
