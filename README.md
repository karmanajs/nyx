# Nyx - Simple Port scanner in Go

**Nyx** is a lightweight port scanner written in Go, designed for quick network reconnaissance and port availability checks.

## Features

- Scan single ports, comma-separated lists, and port ranges (e.g., `80,443,8000-9000`)
- Support for multiple protocols (TCP, UDP, etc.)
- Simple command-line interface
- Fast and efficient scanning
- Clean output format

## Installation

### From Source
```bash
git clone https://github.com/karma-orig/nyx-cli.git
cd nyx-cli
go build -o nyx
```

### From Source
```bash
go install github.com/karma-orig/nyx-cli@latest
```

## Usage

Basic syntax:
```bash
nyx -h <host> [-p <ports>] [-tp <protocol>]
```

### Options

| Flag               | Description                              | Default   |
|--------------------|------------------------------------------|-----------|
| `-h`, `--host`     | Target host (IP or domain)               | *required*|
| `-p`, `--ports`    | Ports to scan (comma-separated or range) | `80,443`  |
| `-tp`, `--type-protocol` | Protocol type                     | `tcp`     |

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

Scan a port range:
```bash
nyx -h example.com -p 8000-9000
```

## Output Example
```bash
Starting Nyx 0.0.1
Scanning example.com (tcp)
80 - open
443 - open
8080 - closed
9000 - closed
```

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

Copyright © 2025 github.com/karmanajs