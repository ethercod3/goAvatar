# goAvatar

![Language](https://img.shields.io/badge/Language-Golang-blue)
![Category](https://img.shields.io/badge/Category-Utility-orange)
![CLI Interface](https://img.shields.io/badge/CLI%20Interface-Yes-yellow)
![API](https://img.shields.io/badge/API-Yes-green)

Go package for quick generation of random pixelized avatars

![examples](https://github.com/ethercod3/goAvatar/blob/main/examples/examples.png?raw=true)

## Example usage

## From CLI

```bash
go install github.com/ethercod3/goAvatar/cmd/goavatar@latest
goavatar -s 500 -d 5
```

For local development:

```bash
go run ./cmd/goavatar -s 500 -d 5
```

### CLI Options

| Short | Long | Required | Description | Default |
|------|------|----------|-------------|---------|
| `-d` | `--dimensions` | Yes | Dimensions of the image | — |
| `-s` | `--size` | Yes | File size in pixels | — |
| `-o` | `--output` | No | Output file | `avatar.png` |

## From code

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

```bash
go test ./...
go build ./...
nu scripts/build.nu
```
