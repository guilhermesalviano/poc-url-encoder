# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

BINARY_NAME=url-encoder

build: deps
	CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' .

test: deps lint
	go test -cover ./...

clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run: build
	./$(BINARY_NAME)

deps:
	$(GOMOD) vendor

#docker-deps:
#	rm -Rf 
#	docker-compose up vendor

#docker-stack:
#	docker network create comm-hub || true
#	docker run --rm --name rabbitmq.local --network comm-hub -d -p 5672:5672 -p 15672:15672 rabbitmq:3.6.10-management-alpine  || true

#docker-run-unit:
#	docker-compose up unit-test

docker-run:
	docker-compose stop url-encoder
	docker-compose up url-encoder

docker-build:
	docker-compose up build

#docker-serve: docker-deps docker-stack docker-run
docker-serve: docker-run