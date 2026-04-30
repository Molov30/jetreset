# Jetreset

[![Release](https://img.shields.io/github/v/release/insigmo/jetreset?style=flat-square)](https://github.com/insigmo/jetreset/releases/latest)
[![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat-square&logo=go)](https://go.dev)


A cross-platform CLI tool for resetting the trial period of JetBrains IDEs.
If you already have licenses in your IDEs, it'll be resetting too.
If you're from Russia, you need to have VPN to start a new trial period.

Supports **IntelliJ IDEA**, **GoLand**, **PyCharm**, **WebStorm**, **CLion**, **PhpStorm**, **Rider**, **DataGrip**, and **RubyMine**.

---

## Installation

### Download binary (recommended)

Download the latest pre-built binary for your platform from the [Releases](https://github.com/insigmo/jetreset/releases/latest) page:

### Linux / macOS

```bash
# Download and run installer
curl -fsSL https://raw.githubusercontent.com/insigmo/jetreset/refs/heads/master/install.sh | bash
```

### Windows (PowerShell)

```powershell
# Download and run installer
irm https://raw.githubusercontent.com/insigmo/jetreset/refs/heads/master/install.ps1 | iex
```

### Run on Linux/MacOS
```bash
./jetreset # Run without flags to reset all detected JetBrains IDEs
```

### Run on Windows
```powershell
.\jetreset.exe
```

### Run scheduler
```bash
./jetreset --run-schedule # Run scheduler for reseting every month
./jetreset --stop         # Stop schedurler
```

### Build from source

Requires Go 1.22+.

```bash
git clone https://github.com/insigmo/jetreset.git
cd jetreset
go build -ldflags="-s -w" -trimpath -o jetreset .
```

---

## Supported IDEs

- IntelliJ IDEA
- GoLand
- PyCharm
- WebStorm
- CLion
- PhpStorm
- Rider
- DataGrip
- RubyMine
