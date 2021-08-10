BINARY_NAME?=plans-service-grpc-poc
BUILD_DIR?=out
GO_LINKER_FLAGS=-ldflags="-s -w"

.PHONY: deps proto-generate proto-generate-php-client run-grpc-server run-http-server run-golang-client run-php-client

proto-generate:
	@protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		--plugin=protoc-gen-grpc=bins/opt/grpc_php_plugin \
         proto/plans_service.proto

proto-generate-php-client:
	@mkdir -p
	@docker run -v `pwd`/proto:/defs namely/protoc-all -f plans_service.proto -l php -o php/

deps:
	@go mod vendor

run-grpc-server:
	@go run cmd/grpc-server.go

run-http-server:
	@go run cmd/http-server.go

run-golang-client:
	@go run cmd/client.go --port 8000

run-php-client:
	@cd php-client && composer dump-autoload --optimize
	@php php-client/client.php
