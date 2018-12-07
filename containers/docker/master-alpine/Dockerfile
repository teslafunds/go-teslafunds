FROM alpine:3.3

RUN \
  apk add --update go git make gcc musl-dev         && \
  git clone https://github.com/teslafunds/go-teslafunds && \
  (cd go-teslafunds && make geth)                     && \
  cp go-teslafunds/build/bin/gtsf /gtsf               && \
  apk del go git make gcc musl-dev                  && \
  rm -rf /go-teslafunds && rm -rf /var/cache/apk/*

EXPOSE 9393
EXPOSE 39993

ENTRYPOINT ["/gtsf"]
