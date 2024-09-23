FROM golang:1.23.1

ENV PROTOBUF_VERSION 4.28.2

RUN apt-get update \
    && apt-get install -y protobuf-compiler unzip --no-install-recommends \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /protoc
COPY . /protoc

## build時のみ使いたいので、runで実行。
RUN chmod +x bin/go_starter && bin/go_starter

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest