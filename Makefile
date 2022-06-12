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

fmt-check:
	@sh scripts/fmt-check.sh
.PHONY: fmt-check

clean:
	@rm -rf $(OUTPUT)
.PHONY: clean
