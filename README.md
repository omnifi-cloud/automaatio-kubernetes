# Automaatio Kubernetes

An installation and management tool for Kubernetes. Automaatio provides the ability to single-click install Kubernetes clusters in various environments for production use cases and development.

## Building

Building the basic binary can be achieved by the simple `go build` command:

```shell
go build -o .bin ./...
```

To support the target environments the following can be used:

```shell
go generate ./... && \
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o .bin/darwin-amd64/ ./... && \
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o .bin/darwin-arm64/ ./... && \
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o .bin/linux-amd64/ ./... && \
CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o .bin/linux-arm/ ./... && \ 
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o .bin/linux-arm64/ ./... && \ 
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o .bin/windows-amd64/ ./... && \ 
CGO_ENABLED=0 GOOS=windows GOARCH=arm go build -o .bin/windows-arm/ ./...
```

## Code quality

Formatting: 

```shell
go fmt ./...
```

Linting:

```shell
go vet ./...
```
