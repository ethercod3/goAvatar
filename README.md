# goAvatar

![Language](https://img.shields.io/badge/Language-Golang-blue)
![Category](https://img.shields.io/badge/Category-Utility-orange)
![CLI Interface](https://img.shields.io/badge/CLI%20Interface-Yes-yellow)
![API](https://img.shields.io/badge/API-Yes-green)

Go package for quick generation of random pixelized avatars

![examples](https://github.com/ethercod3/goAvatar/blob/main/examples/examples.png?raw=true)

# Example usage

## From CLI

```bash
./goAvatar -s 500 -d 5
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
	"os"
	"testing"

	"github.com/ethercod3/goAvatar/avatar"
)

func TestMain(t *testing.T) {
	options := avatar.Options{
		Dimensions: 5,
		FileSizePx: 500,
	}
	img := avatar.Generate(options)
	file, err := os.Create("./avatar.png")
	if err != nil {
		t.Fatal(err)
	}
	png.Encode(file, img)
}

```
