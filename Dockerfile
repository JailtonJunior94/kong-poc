FROM golang:1.21 AS builder

RUN mkdir /go-plugins
WORKDIR /go-plugins

COPY ./plugins/hello ./
COPY ./plugins/auth ./

RUN go mod download && GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o hello .
RUN go mod download && GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o auth .

FROM kong/kong-gateway:3.2.2.5-alpine
USER root

COPY --from=builder ./go-plugins/hello /usr/local/bin/hello
COPY --from=builder ./go-plugins/auth /usr/local/bin/auth

USER kong