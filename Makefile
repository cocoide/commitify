.PHONY: test

test:
	go generate ./...
	go test -v ./internal/service