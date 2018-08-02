FROM golang:1.10-alpine

RUN apk --no-cache add curl git

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN mkdir -p /go/src/github.com/nycdavid/go-incubator

WORKDIR /go/src/github.com/nycdavid/go-incubator

COPY ./ ./

RUN dep ensure
