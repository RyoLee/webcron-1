FROM golang:alpine

ENV TZ Asia/Shanghai
RUN apk update &&\
    apk add git curl jq &&\
    go get github.com/RyoLee/webcron &&\
    mkdir -p /go/src/github.com/RyoLee &&\
    cd /go/src/github.com/RyoLee &&\
    git clone https://github.com/RyoLee/webcron.git &&\
    cd webcron &&\
    go mod init &&\
    go mod tidy &&\
    go build
WORKDIR /go/src/github.com/RyoLee/webcron
CMD ["./webcron"]
