FROM golang:1.13.1-alpine
RUN apk add --no-cache git
RUN apk add --no-cache build-base
RUN go get -u github.com/astaxie/beego
RUN export GO111MODULE=on && go get github.com/beego/bee
WORKDIR /go/src