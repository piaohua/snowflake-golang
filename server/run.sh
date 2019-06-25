#!/bin/bash

set -e

BINFILE=snowflake
SRCFILE=${GOPATH}/src/snowflake-golang/main.go
LOGFILE=./${BINFILE}-server.log

build() {
    export CGO_ENABLED=0
    export GOOS=linux
    export GOARCH=amd64
    export GOPATH=${GOPATH}
    export GO111MODULE=on
    go build -o ${BINFILE} -a -ldflags "-w -s" -installsuffix cgo ${SRCFILE}
}

start() {
    chmod +x ./${BINFILE}
    ./${BINFILE} > ${LOGFILE} 2>&1 &
}

stop() {
    ps aux | grep -v grep | grep "${BINFILE}"
}

test() {
    curl -H "x-snowflake-access-token: snowflake" \
        -H "Content-Type: application/json" \
        -H 'cache-control: no-cache' \
        -X POST http://localhost:8085/v1/id/${1} \
        -i -vv
}

case $1 in
    build)
        build;;
    start)
        start;;
    stop)
        stop;;
    test)
        test $2;;
    *)
        echo "./run.sh build|start|stop|test"
esac
