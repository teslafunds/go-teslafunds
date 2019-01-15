# Build Geth in a stock Go builder container
FROM golang:1.9-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers

ADD . /go-teslafunds
RUN cd /go-teslafunds && make gtsf

# Pull Geth into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /go-teslafunds/build/bin/geth /usr/local/bin/

EXPOSE 9656 9656 42786 42786/udp
ENTRYPOINT ["gtsf"]
