TARGET=service.out

build: proto
	go build -o $(TARGET) src/main.go 

proto:
	 protoc -I api/proto --go_out=plugins=grpc:api/subscription api/proto/subscription.proto

