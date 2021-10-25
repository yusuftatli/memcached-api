FROM golang:latest

WORKDIR /tmp/yemek-sepeti-api

RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates && apk add openssh
RUN apk add build-base
RUN apk add make

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Unit tests
RUN CGO_ENABLED=0 go test -v

# Build the Go app
RUN go build -o ./out/go-sample-app 