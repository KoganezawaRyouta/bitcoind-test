FROM golang:1.10.3-alpine

RUN set -ex && apk add --no-cache make git bash

ENV APP_ROOT /go/src/github.com/KoganezawaRyouta/bitcoind-test
RUN mkdir -p $APP_ROOT
WORKDIR $APP_ROOT
COPY . $APP_ROOT
RUN make setup
