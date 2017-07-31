FROM alpine:3.5

RUN \
<<<<<<< HEAD
  apk add --update go git make gcc musl-dev         && \
  git clone https://github.com/teslafunds/go-teslafunds && \
  (cd go-teslafunds && git checkout develop)          && \
  (cd go-teslafunds && make gtsf)                     && \
  cp go-teslafunds/build/bin/gtsf /gtsf               && \
  apk del go git make gcc musl-dev                  && \
  rm -rf /go-teslafunds && rm -rf /var/cache/apk/*

EXPOSE 3939
EXPOSE 39993
=======
  apk add --update go git make gcc musl-dev ca-certificates && \
  git clone --depth 1 https://github.com/dubaicoin-dbix/go-dubaicoin && \
  (cd go-ethereum && make gdbix) && \
  cp go-ethereum/build/bin/gdbix /gdbix && \
  apk del go git make gcc musl-dev && \
  rm -rf /go-dubaicoin && rm -rf /var/cache/apk/*

EXPOSE 7565
EXPOSE 57955
>>>>>>> 7fdd714... gdbix-update v1.5.0

ENTRYPOINT ["/gtsf"]
