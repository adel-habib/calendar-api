# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /api

COPY init.sh /init.sh

RUN chmod +x /init.sh

EXPOSE 8080

ENTRYPOINT ["sh","/init.sh"]
