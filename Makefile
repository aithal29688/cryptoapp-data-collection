BINARY=crypto-data-collection

VERSION=0.0.1

LDFLAGS=-ldflags ""

GIT_VERSION=$(shell git describe --tags --always --long --dirty --abbrev=7)
GO_PKG=github.com/suaithal/cryptoapp-data-collection

GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)
GOVERSION=1.9.0
OUTPUT_DIR=.
OUTPUT_NAME=${BINARY}
UID=$(shell id -u)
GID=$(shell id -g)


build:
	go build ${LDFLAGS} -o ${BINARY}

fmt:
	gofmt -w ./$*

tests:
	go test github.com/Crypto/cryptoapp-data-collection/

install:
	go install ${LDFLAGS}

dist: clean tests
	GOOS=linux go build ${LDFLAGS} -o ${BINARY}

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: clean install
