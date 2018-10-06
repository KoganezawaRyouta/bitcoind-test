FROM golang:1.10.3-alpine

RUN set -ex && apk add --no-cache make git bash curl unzip

RUN curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.6.1/protoc-3.6.1-linux-x86_64.zip
RUN unzip protoc-3.6.1-linux-x86_64.zip -d protoc3
RUN mv protoc3/bin/* /usr/local/bin/

ENV APP_ROOT /go/src/github.com/KoganezawaRyouta/bitcoind-test
RUN mkdir -p $APP_ROOT

WORKDIR $APP_ROOT
COPY . $APP_ROOT
RUN make setup
