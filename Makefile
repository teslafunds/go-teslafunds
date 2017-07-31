# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

<<<<<<< HEAD
.PHONY: gtsf gtsf-cross evm all test clean
.PHONY: gtsf-linux gtsf-linux-386 gtsf-linux-amd64 gtsf-linux-mips64 gtsf-linux-mips64le
.PHONY: gtsf-linux-arm gtsf-linux-arm-5 gtsf-linux-arm-6 gtsf-linux-arm-7 gtsf-linux-arm64
.PHONY: gtsf-darwin gtsf-darwin-386 gtsf-darwin-amd64
.PHONY: gtsf-windows gtsf-windows-386 gtsf-windows-amd64
.PHONY: gtsf-android gtsf-ios
=======
.PHONY: gdbix android ios gdbix-cross evm all test clean
.PHONY: gdbix-linux gdbix-linux-386 gdbix-linux-amd64 gdbix-linux-mips64 gdbix-linux-mips64le
.PHONY: gdbix-linux-arm gdbix-linux-arm-5 gdbix-linux-arm-6 gdbix-linux-arm-7 gdbix-linux-arm64
.PHONY: gdbix-darwin gdbix-darwin-386 gdbix-darwin-amd64
.PHONY: gdbix-windows gdbix-windows-386 gdbix-windows-amd64
>>>>>>> 7fdd714... gdbix-update v1.5.0

GOBIN = build/bin
GO ?= latest

gtsf:
	build/env.sh go run build/ci.go install ./cmd/gtsf
	@echo "Done building."
	@echo "Run \"$(GOBIN)/gtsf\" to launch gtsf."

evm:
	build/env.sh go run build/ci.go install ./cmd/evm
	@echo "Done building."
	@echo "Run \"$(GOBIN)/evm\" to start the evm."

all:
	build/env.sh go run build/ci.go install

android:
	build/env.sh go run build/ci.go aar --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/gdbix.aar\" to use the library."

ios:
	build/env.sh go run build/ci.go xcode --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/Gdbix.framework\" to use the library."

test: all
	build/env.sh go run build/ci.go test

clean:
	rm -fr build/_workspace/pkg/ $(GOBIN)/*

# Cross Compilation Targets (xgo)

gtsf-cross: gtsf-linux gtsf-darwin gtsf-windows gtsf-android gtsf-ios
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/gtsf-*

gtsf-linux: gtsf-linux-386 gtsf-linux-amd64 gtsf-linux-arm gtsf-linux-mips64 gtsf-linux-mips64le
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/gtsf-linux-*

gtsf-linux-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --dest=$(GOBIN) --targets=linux/386 -v ./cmd/gtsf
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/gtsf-linux-* | grep 386

gtsf-linux-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --dest=$(GOBIN) --targets=linux/amd64 -v ./cmd/gtsf
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gtsf-linux-* | grep amd64

gtsf-linux-arm: gtsf-linux-arm-5 gtsf-linux-arm-6 gtsf-linux-arm-7 gtsf-linux-arm64
	@echo "Linux ARM cross compilation done:"
	@ls -ld $(GOBIN)/gtsf-linux-* | grep arm

gtsf-linux-arm-5:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --dest=$(GOBIN) --targets=linux/arm-5 -v ./cmd/gtsf
	@echo "Linux ARMv5 cross compilation done:"
	@ls -ld $(GOBIN)/gtsf-linux-* | grep arm-5

gtsf-linux-arm-6:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --dest=$(GOBIN) --targets=linux/arm-6 -v ./cmd/gtsf
	@echo "Linux ARMv6 cross compilation done:"
	@ls -ld $(GOBIN)/gtsf-linux-* | grep arm-6

gtsf-linux-arm-7:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --dest=$(GOBIN) --targets=linux/arm-7 -v ./cmd/gtsf
	@echo "Linux ARMv7 cross compilation done:"
	@ls -ld $(GOBIN)/gtsf-linux-* | grep arm-7

gtsf-linux-arm64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --dest=$(GOBIN) --targets=linux/arm64 -v ./cmd/gtsf
	@echo "Linux ARM64 cross compilation done:"
	@ls -ld $(GOBIN)/gtsf-linux-* | grep arm64

gtsf-linux-mips64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --dest=$(GOBIN) --targets=linux/mips64 -v ./cmd/gtsf
	@echo "Linux MIPS64 cross compilation done:"
	@ls -ld $(GOBIN)/gtsf-linux-* | grep mips64

gtsf-linux-mips64le:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --dest=$(GOBIN) --targets=linux/mips64le -v ./cmd/gtsf
	@echo "Linux MIPS64le cross compilation done:"
	@ls -ld $(GOBIN)/gtsf-linux-* | grep mips64le

gtsf-darwin: gtsf-darwin-386 gtsf-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/gtsf-darwin-*

gtsf-darwin-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --dest=$(GOBIN) --targets=darwin/386 -v ./cmd/gtsf
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/gtsf-darwin-* | grep 386

gtsf-darwin-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --dest=$(GOBIN) --targets=darwin/amd64 -v ./cmd/gtsf
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gtsf-darwin-* | grep amd64

gtsf-windows: gtsf-windows-386 gtsf-windows-amd64
	@echo "Windows cross compilation done:"
	@ls -ld $(GOBIN)/gtsf-windows-*

gtsf-windows-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --dest=$(GOBIN) --targets=windows/386 -v ./cmd/gtsf
	@echo "Windows 386 cross compilation done:"
	@ls -ld $(GOBIN)/gtsf-windows-* | grep 386

gtsf-windows-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --dest=$(GOBIN) --targets=windows/amd64 -v ./cmd/gtsf
	@echo "Windows amd64 cross compilation done:"
<<<<<<< HEAD
	@ls -ld $(GOBIN)/gtsf-windows-* | grep amd64

gtsf-android:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --dest=$(GOBIN) --targets=android-21/aar -v ./cmd/gtsf
	@echo "Android cross compilation done:"
	@ls -ld $(GOBIN)/gtsf-android-*

gtsf-ios:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --dest=$(GOBIN) --targets=ios-7.0/framework -v ./cmd/gtsf
	@echo "iOS framework cross compilation done:"
	@ls -ld $(GOBIN)/gtsf-ios-*
=======
	@ls -ld $(GOBIN)/gdbix-windows-* | grep amd64
>>>>>>> 7fdd714... gdbix-update v1.5.0
