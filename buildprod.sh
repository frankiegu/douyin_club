#!/bin/sh

# Build web and other services

cd /opt/go/gopath/src/github.com/Yq2/video_server/api
env GOOS=linux GOARCH=amd64 go build -o ../bin/api

cd /opt/go/gopath/src/github.com/Yq2/video_server/scheduler
env GOOS=linux GOARCH=amd64 go build -o ../bin/scheduler

cd /opt/go/gopath/src/github.com/Yq2/video_server/streamserver
env GOOS=linux GOARCH=amd64 go build -o ../bin/streamserver

cd /opt/go/gopath/src/github.com/Yq2/video_server/web
env GOOS=linux GOARCH=amd64 go build -o ../bin/web
