FROM alpine:3.5

RUN \
<<<<<<< HEAD
  apk add --update go git make gcc musl-dev         && \
  git clone https://github.com/teslafunds/go-teslafunds && \
  (cd go-teslafunds && make geth)                     && \
  cp go-teslafunds/build/bin/gtsf /gtsf               && \
  apk del go git make gcc musl-dev                  && \
  rm -rf /go-teslafunds && rm -rf /var/cache/apk/*
=======
  apk add --update go git make gcc musl-dev ca-certificates && \
  git clone --depth 1 --branch release/1.5 https://github.com/dubaicoin-dbix/go-dubaicoin && \
  (cd go-ethereum && make gdbix) && \
  cp go-ethereum/build/bin/gdbix /gdbix && \
  apk del go git make gcc musl-dev && \
  rm -rf /go-dubaicoin && rm -rf /var/cache/apk/*
>>>>>>> 7fdd714... gdbix-update v1.5.0

EXPOSE 9393
EXPOSE 39993

ENTRYPOINT ["/gtsf"]
