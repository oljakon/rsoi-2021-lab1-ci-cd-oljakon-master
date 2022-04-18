# syntax=docker/dockerfile:1

FROM golang:1.17 AS lang

WORKDIR /app

COPY . .

RUN mkdir app
RUN go get -d ./... && go build -o app ./...

ENV PORT=8080
ENTRYPOINT ["./app/rsoi-lab"]