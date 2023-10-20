FROM golang:1.21 AS builder

RUN mkdir /go-plugins
WORKDIR /go-plugins

COPY ./plugins/hello ./
COPY ./plugins/auth ./

RUN go mod download && go build -o hello .
RUN go mod download && go build -o auth .

FROM kong:3.4.0-ubuntu
USER root

COPY --from=builder ./go-plugins/hello /usr/local/bin/hello
COPY --from=builder ./go-plugins/auth /usr/local/bin/auth

RUN luarocks install kong-oidc && \
    luarocks install kong-jwt2header && \
    luarocks install kong-upstream-jwt

USER kong
ENTRYPOINT ["/docker-entrypoint.sh"]
EXPOSE 8000 8443 8001 8444 8002 8445 8003 8446 8004 8447
STOPSIGNAL SIGQUIT
HEALTHCHECK --interval=10s --timeout=10s --retries=10 CMD kong health
CMD ["kong", "docker-start"]