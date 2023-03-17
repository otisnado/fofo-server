FROM golang:1.20.1-alpine AS build

RUN mkdir /src

WORKDIR /src

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY . /src/

RUN go build .

FROM alpine

RUN adduser -D nepackage

USER nepackage

WORKDIR /home/nepackage

ENV DSN='root:1234@tcp(localhost:3306)/nepackage?parseTime=true&loc=America%2FEl_Salvador'

COPY --from=build /src/nepackage /home/fofo/nepackage

ENTRYPOINT [ "/home/fofo/nepackage" ]