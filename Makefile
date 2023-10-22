

.PHONY: all \
	init \
	dependencies \
	generate_api

init: dependencies generated

generate: generated

generated: generate_api
	mkdir -p generated

dependencies:
	go mod verify
	go mod tidy
	go mod vendor


generate_api: api.yml
	mkdir -p generated/api
	oapi-codegen --package api -generate types $< > generated/api/api-types.gen.go
	oapi-codegen --package api -generate gin,spec $< > generated/api/api-server.gen.go

test:
	go test -short -coverprofile coverage.out -v ./...
