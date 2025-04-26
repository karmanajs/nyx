# Nyx - Simple Port Scanner in Go

**Nyx** is a lightweight port scanner written in Go, designed for quick network reconnaissance and port availability checks.

## Features

- Scan single ports, comma-separated lists, and port ranges (e.g., `80,443,8000-9000`)
- Support for multiple protocols (TCP, UDP, etc.)
- Simple command-line interface
- Fast and efficient scanning
- Clean output format
- JSON export support
- Configurable timeout

## Installation

### From Source (Makefile)
```bash
git clone https://github.com/karmanajs/nyx.git
cd nyx
make build-cli  # Builds binary in ./bin/nyx
```

### From Source (Taskfile)
```bash
task build-cli
```

### Using Go Install
```bash
go install github.com/karmanajs/nyx@latest
```

## Usage

Basic syntax:
```bash
nyx -h <host> [-p <ports>] [-tp <protocol>] [-jf <output.json>]
```

### Options

| Flag               | Description                              | Default   |
|--------------------|------------------------------------------|-----------|
| `-h`, `--host`     | Target host (IP or domain)               | *required*|
| `-p`, `--ports`    | Ports to scan (comma-separated or range) | `80,443`  |
| `-tp`, `--type-protocol` | Protocol type (tcp/udp)            | `tcp`     |
| `-jf`, `--json-file`	| Save results to JSON file	            | `- `      |
| `-to`, `--timeout` 	|Connection timeout (e.g. 500ms, 2s)	| `2s`      |

## Supported Protocols
- tcp, tcp4 (IPv4-only), tcp6 (IPv6-only)
- udp, udp4 (IPv4-only), udp6 (IPv6-only)
- ip, ip4 (IPv4-only), ip6 (IPv6-only)
- unix, unixgram, unixpacket

## Examples

Scan common ports on example.com:
```bash
nyx -h example.com
```

Scan specific ports with UDP protocol:
```bash
nyx -h example.com -p 53,123 -tp udp
```

Scan a port range and save to JSON:
```bash
nyx -h example.com -p 8000-9000 -jf results.json
```

## Output Example
Standard output:
```bash
Starting Nyx 0.0.1
Scanning example.com (tcp)
80 - open
443 - open
8080 - closed
9000 - closed
```
JSON output (when using -jf):
```json
[
  {
    "port": 80,
    "status": "open"
  },
  {
    "port": 443,
    "status": "open"
  },
  {
    "port": 8080,
    "status": "closed"
  }
]
```

## Building and Maintenance
Build with Make
```bash
make build-cli  # Build binary
make clean      # Remove build artifacts
```

Build with Task
```bash
task build-cli  # Build binary
task clean      # Clean up
```

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

Copyright © 2025 github.com/karmanajs