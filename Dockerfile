# ------------------------------------------------------------------------------
# Build Phase
# ------------------------------------------------------------------------------

FROM golang:1.13 AS build

ADD . /go/src/github.com/sosedoff/docker-http-proxy
WORKDIR /go/src/github.com/sosedoff/docker-http-proxy

RUN \
  GOOS=linux \
  GOARCH=amd64 \
  CGO_ENABLED=0 \
  go build -o /docker-http-proxy

# ------------------------------------------------------------------------------
# Package Phase
# ------------------------------------------------------------------------------

FROM alpine:3.6

RUN \
  apk update && \
  apk add --no-cache ca-certificates openssl wget && \
  update-ca-certificates

COPY --from=build /docker-http-proxy /bin/docker-http-proxy

ENV PORT 8080

EXPOSE 8080
CMD ["/bin/docker-http-proxy"]
