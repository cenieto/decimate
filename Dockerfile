FROM golang:latest

# TODO: make entrypoint configurable by env. vars

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
