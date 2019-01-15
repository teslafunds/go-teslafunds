FROM alpine:3.5

RUN \
  apk add --update go git make gcc musl-dev linux-headers ca-certificates && \
  git clone --depth 1 https://github.com/teslafunds/go-teslafunds && \
  (cd go-teslafunds && make gtsf) && \
  cp go-teslafunds/build/bin/gtsf /gtsf && \
  apk del go git make gcc musl-dev linux-headers && \
  rm -rf /go-teslafunds && rm -rf /var/cache/apk/*

EXPOSE 9656
EXPOSE 42786

ENTRYPOINT ["/gtsf"]
