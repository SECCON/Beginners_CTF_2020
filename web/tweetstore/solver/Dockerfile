FROM golang:1.14-alpine as builder

WORKDIR /go/src/app
COPY src .
RUN go get && \
      CGO_ENABLED=0 go build -a -tags "netgo" -installsuffix netgo -ldflags="-s -w -extldflags \"-static\"" -o=/solver

FROM scratch
MAINTAINER Yutaka WATANABE <contact@ytn86.net>

COPY --from=builder /solver /solver
CMD ["/solver"]
