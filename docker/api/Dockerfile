FROM golang:1.23.2-alpine

ENV ROOT=/go/src
ENV CGO_ENABLED 0
ENV GOPATH=

WORKDIR /go/src
COPY go.mod ./go.mod
COPY go.sum ./go.sum

RUN apk update && apk add git
RUN go mod download

WORKDIR /go/src/cmd
CMD [ "go" , "run" , "./main.go" ]
