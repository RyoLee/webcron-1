FROM golang:alpine

ENV TZ Asia/Shanghai
RUN apk update && apk add git && apk add curl
RUN go get github.com/RyoLee/webcron
WORKDIR /go/src/github.com/RyoLee/webcron
RUN  go build
CMD ["./webcron"]
