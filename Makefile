# GITVERSION = $(shell git describe --tags --always --dirty  | tr -d '\n')
# GIT_HASH = $(shell git rev-parse HEAD)

.PHONY: vendor
vendor:
	go mod tidy
	go mod vendor

.PHONY: gofmt
gofmt:
	gofmt -l -s -w ./*.go

.PHONY: run
run:
	go run main.go

.PHONY: build
build:
	CGO_ENABLED=0 go build -v -mod=vendor \
		-o bin/log-generator ./main.go

lint:
	docker run --rm -it -w /sources -v $(shell pwd):/sources golangci/golangci-lint:v1.32.2-alpine golangci-lint run -v