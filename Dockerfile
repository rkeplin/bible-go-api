FROM golang:1.24-bookworm

ENV DB_HOST db
ENV DB_USER test
ENV DB_PASS test
ENV DB_NAME test

WORKDIR /go/src/github.com/rkeplin/bible-go-api

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /go/bin/server
CMD server

EXPOSE 3000
