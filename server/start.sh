#!/bin/bash

cd /go/src/github.com/user/cpt-chewieeee/canvas_world/server

go get -u github.com/kataras/iris
go build server.go

if [ ${APP_ENV} == production ] 
then
  server;
else
  go get github.com/pilu/fresh && fresh;
fi