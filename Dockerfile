FROM golang:1.13 AS Builder

WORKDIR /echoserver

ADD . .

RUN CGO_ENABLED=0 go build -o /tmp/echoserver .

FROM scratch

LABEL maintainer=xuyuanp@gmail.com

COPY --from=builder /tmp/echoserver /echoserver

EXPOSE 8080

ENTRYPOINT ["/echoserver"]
