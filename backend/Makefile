.PHONY: lint test

lint:
	golangci-lint run --timeout=5m --skip-dirs=".*/.*_test.go"

test:
	go test -v ./...