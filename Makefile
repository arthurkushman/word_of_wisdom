export GO111MODULE=on
export GOPROXY=https://proxy.golang.org
export GOSUMDB=off

LOCAL_BIN:=$(CURDIR)/bin

LDFLAGS:=-X 'platform/core/app.Name=circle-api'\
		 -X 'platform/core/app.Version=v1.0.0'

.PHONY: deps
deps:
	$(info Installing binary dependencies...)
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint@v1.58.0

.PHONY: lint
lint:
	$(info Linting go code...)
	go clean --cache
	go mod tidy
	golangci-lint run ./...

.PHONY: build
build:
	$(info Building...)
	go build -ldflags "$(LDFLAGS)" -o ./bin/server ./cmd

.PHONY: run
run:
	$(info Running go code...)
	go run ./...