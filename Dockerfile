FROM golang:1.15.0-alpine

RUN apk upgrade --no-cache && \
    apk add --update && \
    apk add --no-cache --virtual build-dependencies build-base && \
    apk add --no-cache --update \
    git \
    bash \
    curl

RUN go version

COPY . /go/src/github.com/bb3104/sardine
WORKDIR /go/src/github.com/bb3104/sardine

ENV GOBIN /go/bin

RUN go install github.com/gobuffalo/pop/soda
RUN go get github.com/99designs/gqlgen

