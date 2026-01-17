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
