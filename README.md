# goAvatar

[![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?logo=go&logoColor=white)](https://go.dev/)
[![CLI](https://img.shields.io/badge/CLI-goavatar-2F855A)](#cli)
[![Package](https://img.shields.io/badge/API-avatar.Generate-805AD5)](#go-package)

Generate simple random pixel avatars from the command line or from Go code.

![Generated avatar examples](https://github.com/ethercod3/goAvatar/blob/main/examples/examples.png?raw=true)

## Features

- Symmetric pixel-art avatar generation.
- CLI and importable Go package.
- Configurable grid dimensions, output size, and output file.
- Local Taskfile commands for run, test, example, and release builds.

## Quick Start

Install the CLI:

```bash
go install github.com/ethercod3/goAvatar/cmd/goavatar@latest
```

Generate a 500px avatar with a 5x5 grid:

```bash
goavatar -s 500 -d 5 -o avatar.png
```

For local development:

```bash
task run -- -s 500 -d 5 -o avatar.png
```

## CLI

```text
goavatar -s <size> -d <dimensions> [-o <output>]
```

| Short | Long | Required | Description | Default |
| ----- | ---- | -------- | ----------- | ------- |
| `-d` | `--dimensions` | Yes | Avatar grid dimensions | - |
| `-s` | `--size` | Yes | Output image size in pixels | - |
| `-o` | `--output` | No | Output PNG file path | `avatar.png` |

## Go Package

```go
package main

import (
	"image/png"
	"log"
	"os"

	"github.com/ethercod3/goAvatar/avatar"
)

func main() {
	options := avatar.Options{
		Dimensions: 5,
		FileSizePx: 500,
	}

	if err := options.Validate(); err != nil {
		log.Fatal(err)
	}

	img := avatar.Generate(options)
	file, err := os.Create("avatar.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		log.Fatal(err)
	}
}
```

## Development

This project uses [Task](https://taskfile.dev/) for common commands.

| Command | Description |
| ------- | ----------- |
| `task run -- -s 500 -d 5 -o avatar.png` | Run the CLI locally |
| `task example` | Generate `avatar.png` with sample options |
| `task test` | Run Go tests |
| `task build` | Build release binaries into `dist/` |

Equivalent Go commands:

```bash
go test ./...
go build ./...
go run ./cmd/goavatar -s 500 -d 5 -o avatar.png
```

## Release Builds

Build cross-platform binaries:

```bash
nu scripts/build.nu
```

The script writes release artifacts to `dist/`, which is ignored by Git.
