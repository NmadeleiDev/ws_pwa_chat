FROM golang:1.14.2-alpine3.11

ADD . /go/src/application
WORKDIR /go/src/application

RUN go build -o /application -mod=vendor

CMD ["/application"]