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

.PHONY: docker
docker:
	docker build \
		--tag diebietse/log-generator:local .

.PHONY: run-docker
run-docker:
	docker run -v${PWD}/example_logs.txt:/app/example_logs.txt -ti diebietse/log-generator:local

.PHONY: lint
lint:
	docker run --rm -it -w /sources -v $(shell pwd):/sources golangci/golangci-lint:v1.32.2-alpine golangci-lint run -v

.PHONY: test
test:
	go test ./... -cover -race -mod=vendor -v