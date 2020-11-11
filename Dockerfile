ARG GO_VERSION=1.15.1

FROM golang:${GO_VERSION}-buster AS builder
WORKDIR /src

COPY Makefile Makefile
COPY go.mod go.mod
COPY vendor vendor
COPY main.go main.go

RUN make build

FROM scratch
COPY --from=builder /src/bin/log-generator /bin/log-generator

WORKDIR /app
ENTRYPOINT [ "/bin/log-generator" ]
