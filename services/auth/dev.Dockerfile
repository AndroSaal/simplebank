
FROM golang:1.19-alpine

WORKDIR /usr/src/app

COPY ./ ./

RUN apk add --no-cache make \
 && go mod download \
 && go get github.com/githubnemo/CompileDaemon \
 && go install github.com/githubnemo/CompileDaemon

# Пересобирать контейнер каждый раз, когда изменяется исходный код!
ENTRYPOINT CompileDaemon -build="go build -o ./.bin/auth ./cmd/main.go" -command="./.bin/auth"