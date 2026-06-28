package main

import (
	"log"

	"github.com/ethercod3/goAvatar/argparser"
	"github.com/ethercod3/goAvatar/avatar"
)

func main() {
	dimensions, size, output, seed := argparser.Parse()

	options := avatar.Options{
		Dimensions: dimensions,
		FileSizePx: size,
		Seed:       seed,
	}
	if err := avatar.Save(output, options); err != nil {
		log.Fatal(err)
	}
}
