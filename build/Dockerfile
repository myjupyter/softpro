FROM golang:1.13.8

WORKDIR /go/src/softpro

COPY . .

RUN apt update && apt install -y protobuf-compiler
RUN export GO111MODULE=on
RUN export GOBIN="$(go env GOPATH)/bin"
RUN export PATH="$PATH:$(go env GOPATH)/bin"
RUN go get github.com/golang/protobuf/protoc-gen-go

RUN make build
RUN cp application /go/bin

EXPOSE 8001
EXPOSE 8002

CMD ["application", "--dir", "/go/src/softpro/configs"]
