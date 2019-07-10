FROM golang:1.12.6-stretch

ENV DB_HOST db
ENV DB_USER test
ENV DB_PASS test
ENV DB_NAME test

WORKDIR /go/src/github.com/rkeplin/bible-go-api

COPY . .

RUN go get ./
RUN go build -o /go/bin/server
CMD server

EXPOSE 3000
