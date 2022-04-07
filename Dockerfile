FROM golang:1.17-alpine as local

RUN mkdir -p /go/src/github.com/Sakamoto0525/payment-app
WORKDIR /go/src/github.com/Sakamoto0525/payment-app
ADD . /go/src/github.com/Sakamoto0525/payment-app

RUN apk add --update protobuf-dev

RUN go get -u github.com/golang/protobuf/protoc-gen-go

RUN mkdir -p /tmp/protoc && \
    wget "https://github.com/protocolbuffers/protobuf/releases/download/v3.20.0/protoc-3.20.0-linux-x86_64.zip" -O "/tmp/protoc/protobuf.zip"  && \
    cd /tmp/protoc && \
    unzip protobuf.zip && \
    cp /tmp/protoc/bin/protoc /usr/local/bin && \
    chmod go+rx /usr/local/bin/protoc && \
    cd /tmp && \
    rm -r /tmp/protoc

EXPOSE 8000
