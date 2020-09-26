FROM golang/1.14-alpine

RUN apk --no-cache --update-cache add gcc libc-dev linux-headers \
 && go mod download

CMD ["go", "run", "solve.go"]
