#!/bin/sh
export GO111MODULE=auto && export GOPROXY=https://goproxy.cn && go mod tidy
GOOS=linux GOARCH=amd64 go build -o ./bin/go_gateway
docker build -f dockerfile_dashboard -t go-gateteway-dashboard .
docker build -f dockedockrfile_server -t go-gateteway-server .