#!/bin/sh

# Build web UI
cd /opt/go/gopath/src/github.com/Yq2/douyin_club/web
go install
cp /opt/go/gopath/src/github.com/Yq2/douyin_club/bin/web  /opt/go/gopath/src/github.com/Yq2/douyin_club/bin/douyin_club_web_ui/web
cp -R /opt/go/gopath/src/github.com/Yq2/douyin_club/templates /opt/go/gopath/src/github.com/Yq2/douyin_club/bin/douyin_club_web_ui/
