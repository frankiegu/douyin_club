#!/bin/sh

# Build web and other services

cd /opt/go/gopath/src/github.com/Yq2/douyin_club/api
env GOOS=linux GOARCH=amd64 go build -o ../bin/api

cd /opt/go/gopath/src/github.com/Yq2/douyin_club/scheduler
env GOOS=linux GOARCH=amd64 go build -o ../bin/scheduler

cd /opt/go/gopath/src/github.com/Yq2/douyin_club/streamserver
env GOOS=linux GOARCH=amd64 go build -o ../bin/streamserver

cd /opt/go/gopath/src/github.com/Yq2/douyin_club/web
env GOOS=linux GOARCH=amd64 go build -o ../bin/web
