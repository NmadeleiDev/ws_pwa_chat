FROM golang:1.13-alpine

RUN apk add --no-cache build-base

ADD . /go/src/application
WORKDIR /go/src/application

RUN go build -o /application -mod=vendor

CMD ["/application"]