#FROM golang:alpine
FROM ryosetsu/webcron:1.0

ENV TZ Asia/Shanghai
ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn
RUN apk update && apk add jq git curl
# RUN go get github.com/RyoLee/webcron
# WORKDIR /go/src/github.com/RyoLee/webcron
# RUN  go build
# CMD ["./webcron"]
