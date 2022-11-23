APPNAME = keyd

all: build

build: go.sum
	go build -mod=readonly -o build/$(APPNAME) ./cmd/$(APPNAME)

go.sum:
	go mod tidy

.PHONY: build go.sum

lint:
ifeq (,$(shell which golangci-lint))
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2
endif
	golangci-lint run --out-format=tab --timeout=10m

format:
	go fmt ./cmd/*

lint-fix:
ifeq (,$(shell which golangci-lint))
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2
endif
	golangci-lint run --fix --out-format=tab --issues-exit-code=0

.PHONY: lint format lint-fix