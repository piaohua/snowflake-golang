#!/bin/bash

set -e

BINFILE=server
SRCFILE=${GOPATH}/src/github.com/piaohua/snowflake-golang/examples/server.go
IMAGE_NAME='snowflake/server:latest'
CONTAINER_NAME='snowflakeServer'

build() {
    export CGO_ENABLED=0
    export GOOS=linux
    export GOARCH=amd64
    export GOPATH=${GOPATH}
    export GO111MODULE=on
    go build -o ${BINFILE} -a -ldflags "-w -s" -installsuffix cgo ${SRCFILE}
    docker build -t ${IMAGE_NAME} .
}

start() {
    docker run --rm -tid -p 8084:8080 --name ${CONTAINER_NAME} ${IMAGE_NAME}
}

stop() {
    docker stop ${CONTAINER_NAME}
}

case $1 in
    build)
        build;;
    start)
        start;;
    stop)
        stop;;
    *)
        echo "./run.sh build|start|stop"
esac
