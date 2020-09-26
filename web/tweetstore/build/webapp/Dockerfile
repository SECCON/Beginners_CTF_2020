FROM golang:1.14-alpine as builder

WORKDIR /go/src/app
COPY . .

RUN go get && \
      CGO_ENABLED=0 go build -a -tags "netgo" -installsuffix netgo -ldflags="-s -w -extldflags \"-static\"" -o=/webserver

FROM scratch
MAINTAINER Yutaka WATANABE <contact@ytn86.net>

COPY --from=builder /webserver /webserver
COPY ./templates /templates

USER 1000:1000
CMD ["/webserver"]
