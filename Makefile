run:
	go run .

test:
	go test -v ./tests

build:
  go build

docker build:
	docker build -f Dockerfile . -t memcached-api:1.0

docker run:
	docker run -it -p 8080:8080 memcached-api
