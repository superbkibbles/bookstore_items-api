#start from base image 1.12.13
FROM golang:1.14

ENV REPO_URL=github.com/superbkibbles/bookstore_items-api

ENV GOPATH=/app

ENV APP_PATH=$GOPATH/src/$REPO_URL

ENV WORKPATH=$APP_PATH/src
COPY src $WORKPATH
WORKDIR $WORKPATH

RUN go build -o items-api .
# EXPOSE PORT 8003 to the world
EXPOSE 8083

CMD ["./items-api"]
