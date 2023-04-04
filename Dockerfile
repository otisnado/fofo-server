FROM golang:1.20-buster AS build

RUN mkdir /src

WORKDIR /src

ENV CGO_ENABLED=1

RUN apt update && apt upgrade -y

RUN apt install build-essential -y

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY . /src/

RUN go build .

FROM mcr.microsoft.com/openjdk/jdk:17-ubuntu


# Variables required to build
ENV SPRING_CLI=3.0.4
ENV SYSTEM_USER=nepackage
ENV DSN=''


# OS repositories update
RUN apt update && apt upgrade -y

# Installing required packages
RUN apt install unzip zip curl git -y

# Installing Spring CLI
RUN curl https://repo.maven.apache.org/maven2/org/springframework/boot/spring-boot-cli/${SPRING_CLI}/spring-boot-cli-${SPRING_CLI}-bin.zip -o springcli.zip && \
    unzip springcli.zip && \
    mv spring-${SPRING_CLI} /opt/ && \
    ln -s /opt/spring-${SPRING_CLI}/bin/spring /usr/local/bin/spring && \
    echo 'export SPRING_HOME=/opt/spring-${SPRING_CLI}' >> /etc/profile.d/env.sh && \
    rm springcli.zip

# User creation
RUN adduser --shell /bin/bash ${SYSTEM_USER}

# User set up
USER ${SYSTEM_USER}

WORKDIR /home/${SYSTEM_USER}

COPY --chown=${SYSTEM_USER}:${SYSTEM_USER} --from=build /src/nepackage /usr/local/bin/nepackage

EXPOSE 8080

ENTRYPOINT [ "nepackage" ]