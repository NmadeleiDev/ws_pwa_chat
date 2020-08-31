FROM golang:1.14.2-alpine3.11

ADD . /go/src/application
WORKDIR /go/src/application
RUN mkdir files_download && mkdir files_upload

RUN go build -o /application -mod=vendor

RUN mkdir ./files

CMD ["/application"]