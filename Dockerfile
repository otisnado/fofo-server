FROM golang:1.20.1-alpine AS build

RUN mkdir /src

WORKDIR /src

ENV CGO_ENABLED=1

RUN apk add build-base

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY . /src/

RUN go build .

FROM openjdk:17-alpine


# Variables required to build
ENV SPRING_CLI=3.0.4
ENV SYSTEM_USER=nepa
ENV DSN=''


# OS repositories update
RUN apk update

# Installing required packages
RUN apk add curl bash git

# Shell set up
SHELL ["/bin/bash", "-c"]

# User creation
RUN adduser -D -s /bin/bash ${SYSTEM_USER}


# User set up
USER ${SYSTEM_USER}

WORKDIR /home/${SYSTEM_USER}

COPY --chown=${SYSTEM_USER}:${SYSTEM_USER} --from=build /src/setup.sh .

# Installing external dependencies
RUN chmod +x /home/${SYSTEM_USER}/setup.sh && \
/home/${SYSTEM_USER}/setup.sh

COPY --chown=${SYSTEM_USER}:${SYSTEM_USER} --from=build /src/nepackage /usr/local/bin/nepackage

EXPOSE 8080

ENTRYPOINT [ "nepackage" ]