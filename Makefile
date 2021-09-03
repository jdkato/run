LAST_TAG=$(shell git describe --abbrev=0 --tags)
CURR_SHA=$(shell git rev-parse --verify HEAD)

LDFLAGS=-ldflags "-s -w -X main.version=$(LAST_TAG)"

.PHONY: data test lint install rules setup bench compare release

all: build

# make release tag=v0.4.3
release:
	git tag $(tag)
	git push origin $(tag)

# make build os=darwin
# make build os=windows
# make build os=linux
build:
	GOOS=$(os) GOARCH=amd64 go build ${LDFLAGS} -o bin/$(exe) ./cmd/run

arm:
	GOOS=darwin GOARCH=arm64 go build ${LDFLAGS} -o bin/$(exe) ./cmd/run

test:
	go test ./internal/core ./internal/lint ./internal/check ./pkg/glob
	cucumber --format progress
