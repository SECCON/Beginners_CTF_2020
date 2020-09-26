FROM golang:1.14-alpine AS build

WORKDIR /go/src/app
COPY . .

RUN apk --no-cache --update-cache add gcc libc-dev linux-headers \
  && go build -o c4b .


FROM alpine

WORKDIR /go/src/app
COPY . .
COPY --from=build /go/src/app/c4b .

CMD ["/go/src/app/c4b"]
