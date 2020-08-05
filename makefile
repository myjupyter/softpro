run:
	sudo docker container start test happy_darwin
	
	
build:
	go build ./src/main.go 
	./main

proto:
	 protoc -I api/proto --go_out=plugins=grpc:api/subscription api/proto/subscription.proto

stop:
	sudo docker container stop test happy_darwin
