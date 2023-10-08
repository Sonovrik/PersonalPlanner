FROM golang:alpine

# set right timezone
RUN apk add --no-cache tzdata && \
    ln -snf /usr/share/zoneinfo/Europe/Moscow /etc/localtime && \
    echo Europe/Moscow > /etc/timezone

RUN apk add bash

# golangci-lint
RUN wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s

#RUN go get github.com/githubnemo/CompileDaemon
#ENTRYPOINT CompileDaemon -command="./shortify"

COPY run.sh /tmp
ENTRYPOINT ["/bin/bash", "/tmp/run.sh"]