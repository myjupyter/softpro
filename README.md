# Softpro

Test task

## Build Dependencies 
* [Logrus](https://github.com/sirupsen/logrus) - logger
* [gRPC](https://github.com/grpc/grpc) - Google remote procedure calls
* [Viper](https://github.com/spf13/viper) - configurator
* [go-tarantool](https://github.com/tarantool/go-tarantool) - client in Go for Tarantool
* [golangci-lint](https://github.com/golangci/golangci-lint) - Go linters runner 

To generate **.pb.go** files you will need `protobuf-compiler`

For installation instructions, see [Protocol Buffer Compiler Installation](https://grpc.io/docs/protoc-installation/)

## Usage
* `make` - to build
* `make lint` - to run linter
* `make tests` - to run tests
* `make proto` - to generate **.pb.go** files
* `make format` - to run go fmt
* `make run` - launch docker-compose to build containers
* `make stop` - to stop container
* `make clean` - to remove binary
