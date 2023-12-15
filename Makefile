.PHONY: test

test:
	go test -v ./internal/service
proto:
	cd proto/src && protoc --go_out=../gen --go_opt=paths=source_relative \
	--go-grpc_out=../gen --go-grpc_opt=paths=source_relative \
	*.proto