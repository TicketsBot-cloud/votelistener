# Build container
FROM golang:1.24 AS builder

RUN go version

RUN apt-get update && apt-get upgrade -y && apt-get install -y ca-certificates git zlib1g-dev

COPY . /go/src/github.com/TicketsBot-cloud/votelistener
WORKDIR /go/src/github.com/TicketsBot-cloud/votelistener

RUN git submodule update --init --recursive --remote

RUN set -Eeux && \
    go mod download && \
    go mod verify

RUN GOOS=linux GOARCH=amd64 \
    go build \
    -trimpath \
    -o main cmd/votelistener/main.go

# Prod container
FROM ubuntu:latest

RUN apt-get update && apt-get upgrade -y && apt-get install -y ca-certificates curl

COPY --from=builder /go/src/github.com/TicketsBot-cloud/votelistener/main /srv/votelistener/main

RUN chmod +x /srv/votelistener/main

RUN useradd -m container
USER container
WORKDIR /srv/votelistener

CMD ["/srv/votelistener/main"]