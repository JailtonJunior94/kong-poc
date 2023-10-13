FROM golang:1.21 AS builder

RUN mkdir /go-plugins
WORKDIR /go-plugins
COPY ./plugins/hello ./
RUN go mod download && go build -o hello .

FROM kong:3.4.0-ubuntu
COPY --from=builder ./go-plugins/hello /usr/local/bin/hello

USER kong
ENTRYPOINT ["/docker-entrypoint.sh"]
EXPOSE 8000 8443 8001 8444 8002 8445 8003 8446 8004 8447
STOPSIGNAL SIGQUIT
HEALTHCHECK --interval=10s --timeout=10s --retries=10 CMD kong health
CMD ["kong", "docker-start"]