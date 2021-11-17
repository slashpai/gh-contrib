OS ?= $(shell uname -s | tr A-Z a-z)
ARCH ?= $(shell go env GOARCH)
VERSION = v0.1.0
build:
	cd cmd/gh-contrib && GOOS=$(OS) GOARCH=$(ARCH) CGO_ENABLED=0 go build -o ../../gh-contrib
build-local-image:
	docker build . -t gh-contrib:${VERSION}
go-fmt:
	gofmt -s -w .
