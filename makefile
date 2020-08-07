APP_NAME=application
TARGET=$(APP_NAME).out

build: proto
	go build -o $(TARGET) cmd/$(APP_NAME)/main.go 

proto:
	 protoc -I api/proto --go_out=plugins=grpc:api/subscription api/proto/subscription.proto

