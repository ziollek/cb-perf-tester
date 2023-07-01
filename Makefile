BINARY_NAME=cb-perf-tester
BINARY_FILE := ./bin/$(BINARY_NAME)

VERSION ?= local
SCM_COMMIT ?= `git rev-parse HEAD`

.PHONY: build lint lint-deps clean
build:
	@echo ">> building application"
	go build -trimpath -ldflags \
	"-X main.Version=$(VERSION) \
	-X main.SCMCommit=$(SCM_COMMIT)" \
	-o $(BINARY_FILE) \
	./cmd/...


lint-deps:
	@which golangci-lint > /dev/null || \
		(curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.46.1)

lint: lint-deps
	golangci-lint run

clean:
	go clean
	rm $(BINARY_FILE)

