# mcp-hetzner
A Go Model Context Protocol (MCP) server for interacting with the Hetzner Cloud API.

## Build Client
```bash
npm --prefix ./client i
npm --prefix ./client run build
```

## Build Server

```bash
go mod init github.com/MahdadGhasemian/mcp-hetzner
go mod tidy
go build
# go run main.go
```

## Run Client

```bash
node ./client/build/index.js ./mcp-hetzner
```

## Lint
```bash
# install golangci-lint and the run:
golangci-lint run
```

## Inspector
```bash
npx @modelcontextprotocol/inspector
```
