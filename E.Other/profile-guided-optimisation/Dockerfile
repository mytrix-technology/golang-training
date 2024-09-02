FROM arm64v8/golang:1.21 AS builder

ARG GO111MODULE="on"
ARG GOOS="linux"
ARG GOARCH="arm64"
ENV GO111MODULE=${GO111MODULE}
ENV GOOS=${GOOS}
ENV GOARCH=${GOARCH}

RUN mkdir -p /go/src/talaria
COPY . src/talaria
#RUN cd src/talaria && go mod download  && go build && test -x talaria
RUN cd src/talaria && go mod download  && go build -PGO=./talaria.PGO && test -x talaria

RUN mkdir -p /go/src/talaria-plugin
COPY ./talaria-plugin  src/talaria-plugin
RUN cd src/talaria-plugin && make plugin && test -f talaria-plugin.so
FROM arm64v8/debian:latest AS base

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /root/
ARG GO_BINARY=talaria
COPY  --from=builder /go/src/talaria/${GO_BINARY} .
COPY  --from=builder /go/src/talaria-plugin/talaria-plugin.so .

ADD entrypoint.sh .
RUN mkdir /etc/talaria/ && chmod +x /root/${GO_BINARY} /root/entrypoint.sh
ENV TALARIA_RC=/etc/talaria/talaria.rc
EXPOSE 8027
ENTRYPOINT ["/root/entrypoint.sh"]