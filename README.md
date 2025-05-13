# mcp-hetzner
A Go Model Context Protocol (MCP) server for interacting with the Hetzner Cloud API.

## 🛠 Build Client
```bash
npm --prefix ./client i
npm --prefix ./client run build
```

## 🖥 Build Server

```bash
go mod init github.com/MahdadGhasemian/mcp-hetzner
go mod tidy
go build
# go run main.go
```

## 🚀 Run Client

```bash
node ./client/build/index.js ./mcp-hetzner
```

## ✅ Lint
```bash
# install golangci-lint and the run:
golangci-lint run
```

## 🔍 Inspector
```bash
npx @modelcontextprotocol/inspector
```

## 🗺 Roadmap
- [x] Implement all **GET** and **LIST** operations for:
  - [x] Certificates
  - [x] SSH Keys
  - [x] Locations
  - [x] Datacenters
  - [x] Firewall
  - [x] Floating IPs
  - [x] Servers
  - [x] Images
  - [x] ISOs
  - [x] Placement Groups
  - [x] Primary IPs
  - [x] Server Typs
  - [x] Load Balancers
  - [x] Load Balancer Types
  - [x] Networks
  - [x] Valumes
  - [x] Pricing

- [ ] Add a configuration flag or setting to:
  - [ ] Enable **read-only mode** (GET/LIST only)
  - [ ] Enable **read-write mode** (GET/LIST + CREATE/UPDATE/DELETE)

- [ ] Implement **write operations** **create/update**:
  - [ ] Certificates
  - [ ] SSH Keys
  - [ ] Firewall
  - [ ] Floating IPs
  - [ ] Servers
  - [ ] Images
  - [ ] Placement Groups
  - [ ] Primary IPs
  - [ ] Load Balancers
  - [ ] Networks
  - [ ] Valumes

- [ ] Add **delete capabilities** for supported resources