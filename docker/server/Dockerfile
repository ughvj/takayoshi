FROM golang:1.23-alpine

WORKDIR /go/src/takamori

COPY ./server .

RUN apk upgrade --update && \
    apk --no-cache add git gcc musl-dev

RUN go get github.com/labstack/echo/v4
RUN go get github.com/labstack/echo/middleware
RUN go get -u github.com/go-sql-driver/mysql
