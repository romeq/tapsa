SHELL 		= /bin/bash
GO			= go
BINARY		= usva
identifier	= usva

.PHONY: all lint test

install: build setup

build: 
	$(GO) build -o $(BINARY)

test:
	go test ./...

lint:
	golangci-lint run ./...

format:
	go fmt ./...

clean:
	rm -f $(BINARY)

