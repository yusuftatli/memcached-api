run:
	go run .

test:
	go test -v ./tests

build:
  go build

docker build:
	docker build -f Dockerfile . -t memcash-api:1.0