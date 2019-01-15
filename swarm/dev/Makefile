.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/teslafunds/go-teslafunds/cmd/swarm
	go build -o bin/geth     github.com/teslafunds/go-teslafunds/cmd/geth
	go build -o bin/bootnode github.com/teslafunds/go-teslafunds/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/teslafunds/go-teslafunds/swarm/...
