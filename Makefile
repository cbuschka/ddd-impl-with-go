TOP_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

build:
	go build -o build/orderex ./cli/orderex.go
