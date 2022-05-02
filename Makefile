GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)

.PHONY: wire
# wire
wire:
	cd dingtalk && wire