FROM golang:alpine

ENV TZ Asia/Shanghai
RUN apk update && apk add git && apk add curl && apk add jq
RUN go get github.com/RyoLee/webcron
WORKDIR /go/src/github.com/RyoLee
RUN git clone https://github.com/RyoLee/webcron.git && go mod init && go mod tidy && go build
CMD ["./webcron"]
