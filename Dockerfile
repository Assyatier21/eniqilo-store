FROM golang:1.19

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN apt-get update && apt-get install

RUN go build -o /app_bin ./cmd/main.go

EXPOSE 8080

ENTRYPOINT ["migration.sh"]

CMD ["/app_bin"]