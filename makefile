TARGET=application

all:
	go build -o $(TARGET) cmd/$(TARGET)/main.go

lint:
	golangci-lint run -E golint

proto:
	protoc -I api/proto --go_out=plugins=grpc:api/subscription api/proto/subscription.proto

format:
	go fmt ./...
	
tests:
	$(MAKE) -C test

clean:
	rm $(TARGET)

run:
	docker-compose up -d

stop:
	docker-compose stop
