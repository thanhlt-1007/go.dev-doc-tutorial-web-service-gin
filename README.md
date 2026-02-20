# Tutorial: Developing a RESTful API with Go and Gin

- Reference: https://go.dev/doc/tutorial/web-service-gin

## Installing / Getting started

### go

```shell
gvm install go.26.0
gvm use go.26.0
```

### go packages

```sh
go mod tidy
```

### golangci-lint

```sh
go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.10.1
```

## Commands

### gofumpt

```sh
gofumpt -l -w .
```

### golangci-lint

```sh
golangci-lint run ./...
```

### run

```sh
go run main.go
```
