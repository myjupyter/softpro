TARGET=application
DATA_DIR=tarantool

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

run:
	mkdir -p $(DATA_DIR)
	docker-compose up -d

stop:
	docker-compose stop

clean:
	rm $(TARGET)
