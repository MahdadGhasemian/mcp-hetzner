# mcp-hetzner
A Go Model Context Protocol (MCP) server for interacting with the Hetzner Cloud API.

---

## 🛠 Build Client
```bash
npm --prefix ./client i
npm --prefix ./client run build
```

## 🖥 Build Server
```bash
go mod init github.com/MahdadGhasemian/mcp-hetzner-go
go mod tidy
go build -o mcphetzner
# go run main.go
```

## 🚀 Run Client
```bash
node ./client/build/index.js ./mcp-hetzner
```

---

## ⚠️ Usage Restrictions

The server supports two operation modes, controlled by the configuration:

- **Read-Only mode**: Only allows **GET** and **LIST** operations.
- **Read-Write mode**: Allows **GET**, **LIST**, and **CREATE/UPDATE/DELETE** operations.

**By default, the server starts in read-only mode.**  
To enable write operations (such as creating, updating, or deleting resources), you must explicitly set the configuration flag to enable read-write mode.

> **Warning**: Enabling write mode allows the client to make changes to your Hetzner Cloud resources.  
> Ensure you understand the implications and have proper access controls in place.

### Switching Modes

Edit your configuration or pass the relevant environment variable/flag at launch:

- For **read-only** (default; safe for inspection and monitoring):
    ```
    ./mcphetzner --mode=read_only
    ```
- For **read-write** (use with caution!):
    ```
    ./mcphetzner --mode=read_write
    ```

---

## ✅ Lint
```bash
# install golangci-lint and then run:
golangci-lint run
```

## 🔍 Inspector
```bash
npx @modelcontextprotocol/inspector
```

---

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
  - [x] Server Types
  - [x] Load Balancers
  - [x] Load Balancer Types
  - [x] Networks
  - [x] Volumes
  - [x] Pricing

- [x] Add a configuration flag or setting to:
  - [x] Enable **read_only mode** (GET/LIST only)
  - [x] Enable **read_write mode** (GET/LIST + CREATE/UPDATE/DELETE)

- [ ] Implement **write operations** (**create/update**):
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
  - [ ] Volumes

- [ ] Add **delete capabilities** for supported resources

---