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

	if dimensions == 0 || size == 0 {
		log.Println("Dimensions and size must be specified")
		return
	}

	img := avatar.Generate(avatar.Options{
		Dimensions: dimensions,
		FileSizePx: size,
	})

	file, err := os.Create(output)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	png.Encode(file, img)
}
