FROM golang:1.11-alpine

ENV GO111MODULE=on

WORKDIR /go/src/github.com/emman27/hackerrank

RUN apk add --update --no-cache alpine-sdk

ADD go.mod go.sum ./

RUN go mod download

COPY . .

RUN go test ./...
