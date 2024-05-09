FROM golang:1.120

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN apt-get update && apt-get install

RUN go build -o /app_bin ./cmd/main.go

EXPOSE 8080

CMD ["/app_bin"]