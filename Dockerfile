# syntax=docker/dockerfile:experimental

# Build Image
ARG GO_VERSION=1.14
FROM golang:${GO_VERSION}-alpine AS builder
ENV GO111MODULE=on
RUN apk add --no-cache --update \
        openssh-client \
        git \
        ca-certificates \
        build-base
WORKDIR /go/src/github.com/eahrend/sql-scaler
COPY ./ ./
RUN go mod download
WORKDIR /go/src/github.com/eahrend/sql-scaler
RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o /app .

# Application layer
FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN update-ca-certificates
RUN apk add bash
RUN mkdir /app
COPY --from=builder /app /app
WORKDIR /app
CMD ["./app"]
