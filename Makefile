# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: gtsf android ios gtsf-cross swarm evm all test clean
.PHONY: gtsf-linux gtsf-linux-386 gtsf-linux-amd64 gtsf-linux-mips64 gtsf-linux-mips64le
.PHONY: gtsf-linux-arm gtsf-linux-arm-5 gtsf-linux-arm-6 gtsf-linux-arm-7 gtsf-linux-arm64
.PHONY: gtsf-darwin gtsf-darwin-386 gtsf-darwin-amd64
.PHONY: gtsf-windows gtsf-windows-386 gtsf-windows-amd64

GOBIN = $(shell pwd)/build/bin
GO ?= latest

gtsf:
	build/env.sh go run build/ci.go install ./cmd/gtsf
	@echo "Done building."
	@echo "Run \"$(GOBIN)/gtsf\" to launch gtsf."

swarm:
	build/env.sh go run build/ci.go install ./cmd/swarm
	@echo "Done building."
	@echo "Run \"$(GOBIN)/swarm\" to launch swarm."

all:
	build/env.sh go run build/ci.go install

android:
	build/env.sh go run build/ci.go aar --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/gtsf.aar\" to use the library."

ios:
	build/env.sh go run build/ci.go xcode --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/Gtsf.framework\" to use the library."

test: all
	build/env.sh go run build/ci.go test

clean:
	rm -fr build/_workspace/pkg/ $(GOBIN)/*

# The devtools target installs tools required for 'go generate'.
# You need to put $GOBIN (or $GOPATH/bin) in your PATH to use 'go generate'.

devtools:
	env GOBIN= go get -u golang.org/x/tools/cmd/stringer
	env GOBIN= go get -u github.com/jteeuwen/go-bindata/go-bindata
	env GOBIN= go get -u github.com/fjl/gencodec
	env GOBIN= go install ./cmd/abigen

# Cross Compilation Targets (xgo)

gtsf-cross: gtsf-linux gtsf-darwin gtsf-windows gtsf-android gtsf-ios
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/gtsf-*

gtsf-linux: gtsf-linux-386 gtsf-linux-amd64 gtsf-linux-arm gtsf-linux-mips64 gtsf-linux-mips64le
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/gtsf-linux-*

gtsf-linux-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/386 -v ./cmd/gtsf
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/gtsf-linux-* | grep 386

gtsf-linux-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/amd64 -v ./cmd/gtsf
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gtsf-linux-* | grep amd64

gtsf-linux-arm: gtsf-linux-arm-5 gtsf-linux-arm-6 gtsf-linux-arm-7 gtsf-linux-arm64
	@echo "Linux ARM cross compilation done:"
	@ls -ld $(GOBIN)/gtsf-linux-* | grep arm

gtsf-linux-arm-5:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-5 -v ./cmd/gtsf
	@echo "Linux ARMv5 cross compilation done:"
	@ls -ld $(GOBIN)/gtsf-linux-* | grep arm-5

gtsf-linux-arm-6:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-6 -v ./cmd/gtsf
	@echo "Linux ARMv6 cross compilation done:"
	@ls -ld $(GOBIN)/gtsf-linux-* | grep arm-6

gtsf-linux-arm-7:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-7 -v ./cmd/gtsf
	@echo "Linux ARMv7 cross compilation done:"
	@ls -ld $(GOBIN)/gtsf-linux-* | grep arm-7

gtsf-linux-arm64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm64 -v ./cmd/gtsf
	@echo "Linux ARM64 cross compilation done:"
	@ls -ld $(GOBIN)/gtsf-linux-* | grep arm64

gtsf-linux-mips:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips --ldflags '-extldflags "-static"' -v ./cmd/gtsf
	@echo "Linux MIPS cross compilation done:"
	@ls -ld $(GOBIN)/gtsf-linux-* | grep mips

gtsf-linux-mipsle:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mipsle --ldflags '-extldflags "-static"' -v ./cmd/gtsf
	@echo "Linux MIPSle cross compilation done:"
	@ls -ld $(GOBIN)/gtsf-linux-* | grep mipsle

gtsf-linux-mips64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64 --ldflags '-extldflags "-static"' -v ./cmd/gtsf
	@echo "Linux MIPS64 cross compilation done:"
	@ls -ld $(GOBIN)/gtsf-linux-* | grep mips64

gtsf-linux-mips64le:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64le --ldflags '-extldflags "-static"' -v ./cmd/gtsf
	@echo "Linux MIPS64le cross compilation done:"
	@ls -ld $(GOBIN)/gtsf-linux-* | grep mips64le

gtsf-darwin: gtsf-darwin-386 gtsf-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/gtsf-darwin-*

gtsf-darwin-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/386 -v ./cmd/gtsf
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/gtsf-darwin-* | grep 386

gtsf-darwin-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/amd64 -v ./cmd/gtsf
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gtsf-darwin-* | grep amd64

gtsf-windows: gtsf-windows-386 gtsf-windows-amd64
	@echo "Windows cross compilation done:"
	@ls -ld $(GOBIN)/gtsf-windows-*

gtsf-windows-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/386 -v ./cmd/gtsf
	@echo "Windows 386 cross compilation done:"
	@ls -ld $(GOBIN)/gtsf-windows-* | grep 386

gtsf-windows-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/amd64 -v ./cmd/gtsf
	@echo "Windows amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gtsf-windows-* | grep amd64
