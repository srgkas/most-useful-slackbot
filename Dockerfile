FROM golang:1.16 as builder
WORKDIR go/src/slackbot
COPY . .
RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' -a \
    -o /go/bin/slackbot .

FROM alpine

WORKDIR /app

COPY --from=builder /go/bin/slackbot /slackbot

EXPOSE 80
EXPOSE 8000

VOLUME /app

ENTRYPOINT ["/slackbot"]