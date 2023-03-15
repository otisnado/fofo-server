FROM golang:1.20.1-alpine AS build

RUN mkdir /src

WORKDIR /src

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY . /src/

RUN go build .

FROM alpine

RUN adduser -D fofo

USER fofo

WORKDIR /home/fofo

ENV BDHOST=mysql DBPORT=3306 DBNAME=fofo DBUSER=root DBPASS=1234

COPY --from=build /src/nepackage /home/fofo/nepackage

ENTRYPOINT [ "/home/fofo/nepackage" ]