FROM golang:latest

WORKDIR /app

COPY . /app

CMD go run cmd/app/main.go