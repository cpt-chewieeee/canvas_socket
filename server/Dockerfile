FROM golang

# if left blank server will run with dev settings
# production: docker build ./server --build-args app_env=production
ARG app_en
ENV APP_ENV $app_env

# WORKDIR /go/src/github.com/user/cpt-chewieeee/canvas_world/server
COPY ./start.sh /usr/local/bin
RUN chmod +x /usr/local/bin/start.sh