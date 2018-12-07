FROM ubuntu:wily
MAINTAINER chrisfranko


ENV DEBIAN_FRONTEND noninteractive

# Usual update / upgrade
RUN apt-get update
RUN apt-get upgrade -q -y
RUN apt-get dist-upgrade -q -y

# Let our containers upgrade themselves
RUN apt-get install -q -y unattended-upgrades

# Install Teslafunds
RUN apt-get install -q -y curl git mercurial binutils bison gcc make libgmp3-dev build-essential

# Install Go
RUN \
  mkdir -p /goroot && \
  curl https://storage.googleapis.com/golang/go1.4.2.linux-amd64.tar.gz | tar xvzf - -C /goroot --strip-components=1

# Set environment variables.
ENV GOROOT /goroot
ENV GOPATH /gopath
ENV PATH $GOROOT/bin:$GOPATH/bin:$PATH

RUN git clone http://www.github.com/teslafunds/go-teslafunds.git
RUN cd go-teslafunds && make gtsf

RUN cp build/bin/gtsf /usr/bin/gtsf

EXPOSE 3939
EXPOSE 39993

ENTRYPOINT ["/usr/bin/gtsf"]
