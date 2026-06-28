package argparser

import (
	"flag"
)

func Parse() (int, int, string, int64) {
	var dimensions int
	var size int
	var output string
	var seed int64

	flag.IntVar(&dimensions, "d", 0, "Dimensions of the image")
	flag.IntVar(&dimensions, "dimensions", 0, "Dimensions of the image")

	flag.IntVar(&size, "s", 0, "Size in pixels")
	flag.IntVar(&size, "size", 0, "Size in pixels")

	flag.StringVar(&output, "o", "avatar.png", "output file")
	flag.StringVar(&output, "output", "avatar.png", "output file")

	flag.Int64Var(&seed, "seed", 0, "random seed; 0 uses a random seed")

	flag.Parse()

	return dimensions, size, output, seed
}
