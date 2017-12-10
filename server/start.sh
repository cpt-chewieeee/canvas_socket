#!/bin/bash

cd /go/src/github.com/user/cpt-chewieeee/canvas_world/server

go get github.com/julienschmidt/httprouter
go get github.com/graarh/golang-socketio
go get github.com/lib/pq
go build server.go

if [ ${APP_ENV} == production ] 
then
  server;
else
  go get github.com/pilu/fresh && fresh;
fi