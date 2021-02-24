NAME = ads-node-module
BIN_DIR ?= ../bin
VERSION ?= $(shell git describe --match=NeVeRmAtCh --always --abbrev=40 --dirty)
GO_LDFLAGS = -ldflags "-X ads-node-module/internal/message/version.Version=$(VERSION)"

build:
	GOOS=linux GOARCH=amd64 go build $(GO_LDFLAGS) -o ${BIN_DIR}/${NAME} main.go

test:
	go test ./...
clean:
	rm -f ${BIN_DIR}/${NAME}

.PHONY: clean
