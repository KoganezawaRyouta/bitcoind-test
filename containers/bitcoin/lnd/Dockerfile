FROM golang:1.10.3-alpine

RUN set -ex && apk add --no-cache make git bash curl unzip

RUN go get -u github.com/golang/dep/cmd/dep
RUN go get -d github.com/lightningnetwork/lnd

WORKDIR ${GOPATH}/src/github.com/lightningnetwork/lnd
RUN make
RUN make install

