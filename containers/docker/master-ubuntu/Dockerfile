FROM ubuntu:xenial

RUN \
  apt-get update && apt-get upgrade -q -y && \
  apt-get install -y --no-install-recommends golang git make gcc libc-dev ca-certificates && \
  git clone --depth 1 --branch release/1.7 https://github.com/teslafunds/go-teslafunds && \
  (cd go-teslafunds && make gtsf) && \
  cp go-teslafunds/build/bin/gtsf /gtsf && \

  apt-get remove -y golang git make gcc libc-dev && apt autoremove -y && apt-get clean && \
  rm -rf /go-teslafunds

EXPOSE 4949
EXPOSE 59997

ENTRYPOINT ["/gtsf"]
