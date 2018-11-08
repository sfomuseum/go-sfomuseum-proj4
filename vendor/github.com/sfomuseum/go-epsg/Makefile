CWD=$(shell pwd)
GOPATH := $(CWD)

prep:
	if test -d pkg; then rm -rf pkg; fi

self:   prep rmdeps
	if test -d src; then rm -rf src; fi
	mkdir -p src/github.com/sfomuseum/go-epsg
	cp -r *.go src/github.com/sfomuseum/go-epsg/
	if test -d vendor; then cp -r vendor/* src/; fi

rmdeps:
	if test -d src; then rm -rf src; fi 

build:	fmt bin

deps:
	# @GOPATH=$(GOPATH) go get -u "github.com/go-spatial/proj"

vendor-deps: rmdeps deps
	if test ! -d vendor; then mkdir vendor; fi
	if test -d vendor; then rm -rf vendor; fi
	cp -r src vendor
	find vendor -name '.git' -print -type d -exec rm -rf {} +
	rm -rf src

defs:	bin
	bin/mk-definitions > definitions.go

fmt:
	go fmt *.go
	go fmt cmd/*.go

bin: 	self
	rm -rf bin/*
	@GOPATH=$(GOPATH) go build -o bin/epsg cmd/epsg.go
	@GOPATH=$(GOPATH) go build -o bin/mk-definitions cmd/mk-definitions.go
