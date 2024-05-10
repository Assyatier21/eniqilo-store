FROM golang:1.22.1-alpine AS builder

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o /app_bin ./cmd/main.go

EXPOSE 8080

CMD ["/app_bin"]