FROM golang:1.13.7

WORKDIR /go/src/cookify

COPY go.mod .
COPY go.sum .
RUN go get -d -v ./...

COPY . .

RUN go build -o cookify ./cmd/webapi/main.go

EXPOSE 3000

CMD ["./cookify"]
