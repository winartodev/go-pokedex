FROM golang:1.17-alpine AS builder

RUN apk update && apk upgrade && \
    apk --update add git make curl

RUN mkdir /go-pokedex

ADD . /go-pokedex

WORKDIR /go-pokedex

COPY env.sample .env

RUN go mod download

RUN go build app/main.go

EXPOSE 8080

ENTRYPOINT ["/go-pokedex/main"]