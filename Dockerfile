FROM golang:1.15-alpine AS GO_BUILD
WORKDIR /tmp/memcash-api

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Unit tests
RUN CGO_ENABLED=0 go test -v

# Build the Go app
RUN go build -o ./out/memcash-api
EXPOSE 8080
