# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY ./ ./
RUN go mod download

RUN go get -u github.com/go-sql-driver/mysql

RUN go build -o /docker-gs-ping main.go

EXPOSE 9090

CMD [ "/docker-gs-ping" ]