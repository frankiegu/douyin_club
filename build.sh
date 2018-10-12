#!/bin/sh

# Build web UI
cd /opt/go/gopath/src/github.com/Yq2/video_server/web
go install
cp /opt/go/gopath/src/github.com/Yq2/video_server/bin/web  /opt/go/gopath/src/github.com/Yq2/video_server/bin/video_server_web_ui/web
cp -R /opt/go/gopath/src/github.com/Yq2/video_server/templates /opt/go/gopath/src/github.com/Yq2/video_server/bin/video_server_web_ui/
