FROM golang:1.13.8

RUN git clone http://github.com/myjupyter/softpro /go/src/softpro
WORKDIR /go/src/softpro

RUN apt update && apt install -y protobuf-compiler
RUN export GO111MODULE=on
RUN export GOBIN="$(go env GOPATH)/bin"
RUN export PATH="$PATH:$(go env GOPATH)/bin"
RUN go get github.com/golang/protobuf/protoc-gen-go

RUN make build
RUN cp application.out /go/bin

EXPOSE 8001
EXPOSE 8002

CMD ["application.out", "--dir", "/go/src/softpro/configs"]
