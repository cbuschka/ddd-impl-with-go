TOP_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

test:
	go test -cover -race -coverprofile=coverage.txt -covermode=atomic ./order/... ./cli/...

build:
	go build -o build/orderex ./cli/orderex.go
