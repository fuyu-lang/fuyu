GO ?= go
OUTPUT = ./output

all: build
.PHONY: all

build:
	@mkdir -p $(OUTPUT)
	@$(GO) build -o $(OUTPUT) ./...
.PHONY: build

test:
	@$(GO) test ./...
.PHONY: test

fmt:
	@$(GO) fmt ./...
.PHONY: fmt

clean:
	@rm -rf $(OUTPUT)
.PHONY: clean
