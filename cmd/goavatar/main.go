package main

import (
	"image/png"
	"log"
	"os"

	"github.com/ethercod3/goAvatar/argparser"
	"github.com/ethercod3/goAvatar/avatar"
)

func main() {
	dimensions, size, output := argparser.Parse()

	options := avatar.Options{
		Dimensions: dimensions,
		FileSizePx: size,
	}
	if err := options.Validate(); err != nil {
		log.Fatal(err)
	}

	img := avatar.Generate(options)
	file, err := os.Create(output)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		log.Fatal(err)
	}
}
