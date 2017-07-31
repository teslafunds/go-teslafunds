FROM ubuntu:xenial

RUN \
  apt-get update && apt-get upgrade -q -y && \
  apt-get install -y --no-install-recommends golang git make gcc libc-dev ca-certificates && \
  git clone --depth 1 https://github.com/dubaicoin-dbix/go-dubaicoin && \
  (cd go-ethereum && make gdbix) && \
  cp go-ethereum/build/bin/gdbix /gdbix && \
  apt-get remove -y golang git make gcc libc-dev && apt autoremove -y && apt-get clean && \
  rm -rf /go-dubaicoin

<<<<<<< HEAD
RUN apt-get update && \
    apt-get upgrade -q -y && \
    apt-get dist-upgrade -q -y && \
    apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 923F6CA9 && \
    echo "deb http://ppa.launchpad.net/ethereum/ethereum-dev/ubuntu wily main" | tee -a /etc/apt/sources.list.d/ethereum.list && \
    apt-get update && \
    apt-get install -q -y gtsf

EXPOSE 8545
EXPOSE 30303

ENTRYPOINT ["/usr/bin/gtsf"]
=======
EXPOSE 7565
EXPOSE 57955

ENTRYPOINT ["/gdbix"]
>>>>>>> 7fdd714... gdbix-update v1.5.0
