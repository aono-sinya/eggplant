FROM golang:1.23.1-alpine

RUN apk update && \
    apk --no-cache add git && \
    apk add bash

RUN mkdir /go/src/app
WORKDIR /go/src/app

#CMD ["go", "run", "main.go"]